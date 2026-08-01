const std = @import("std");
const mem = std.mem;
const fmt = std.fmt;
const fs = std.fs;

// Print

pub fn print(comptime element: []const u8, args: anytype) void {
    std.debug.print(element, args);
}

pub fn printStr(str: []const u8) void {
    std.debug.print("{s}\n", .{str});
}

pub fn printAny(value: anytype) void {
    std.debug.print("{any}\n", .{value});
}

pub fn printMap(map: anytype) void {
    var it = map.iterator();
    while (it.next()) |entry| {
        std.debug.print("{any}: {any}\n", .{ entry.key_ptr.*, entry.value_ptr.* });
    }
}

// Types

// Parses a bool from a string
// true:  "1", "true", "TRUE", "True", "yes"
// false: "0", "false", "FALSE", "False", "no"
pub fn stringToBool(str: []const u8) error{InvalidBool}!bool {
    if (str.len == 0 or str.len > 5) return error.InvalidBool;

    // Single-char fast path
    if (str.len == 1) return switch (str[0]) {
        '1' => true,
        '0' => false,
        else => error.InvalidBool,
    };

    var lower_buf: [5]u8 = undefined;
    for (str, 0..) |c, i| lower_buf[i] = std.ascii.toLower(c);
    const lower = lower_buf[0..str.len];

    if (mem.eql(u8, lower, "true") or mem.eql(u8, lower, "yes")) return true;
    if (mem.eql(u8, lower, "false") or mem.eql(u8, lower, "no")) return false;
    return error.InvalidBool;
}

// Works for any integer or float type
// Returns a slice into `buf`. Returns empty slice if buf is too small
pub fn numberToString(buf: []u8, value: anytype) []const u8 {
    const T = @TypeOf(value);
    return switch (@typeInfo(T)) {
        .int, .comptime_int => fmt.bufPrint(buf, "{d}", .{value}) catch buf[0..0],
        .float, .comptime_float => fmt.bufPrint(buf, "{d}", .{value}) catch buf[0..0],
        else => @compileError("numberToString only supports integer and float types"),
    };
}

// Slice

// caller return memory
pub fn sliceToString(allocator: mem.Allocator, comptime T: type, slice: []const T, delimiter: []const u8) ![]u8 {
    if (slice.len == 0) return try allocator.alloc(u8, 0);

    var list = std.ArrayList(u8).init(allocator);
    errdefer list.deinit();

    for (slice, 0..) |item, i| {
        if (i > 0) try list.appendSlice(delimiter);
        const formatted = try fmt.allocPrint(allocator, "{any}", .{item});
        defer allocator.free(formatted);
        try list.appendSlice(formatted);
    }

    return list.toOwnedSlice();
}

// Checks whether a slice contains a given value.
pub fn contains(comptime T: type, slice: []const T, value: T) bool {
    return mem.indexOfScalar(T, slice, value) != null;
}

// Same as contains, but for string slices ([]const u8).
pub fn containsString(slice: []const []const u8, value: []const u8) bool {
    for (slice) |item| {
        if (mem.eql(u8, item, value)) return true;
    }
    return false;
}

// Files

pub const LineWriteError = error{ LineOutOfRange } || fs.File.OpenError || fs.File.WriteError || mem.Allocator.Error || fs.Dir.ReadFileAllocError;

pub const Lines = struct {
    items: [][]const u8,
    _backing: []u8,
    allocator: mem.Allocator,

    pub fn deinit(self: *Lines) void {
        self.allocator.free(self.items);
        self.allocator.free(self._backing);
    }

    pub fn len(self: Lines) usize {
        return self.items.len;
    }

    pub fn get(self: Lines, index: usize) []const u8 {
        return self.items[index];
    }
};


// Reads entire file into a single string
pub fn readFile(allocator: mem.Allocator, file_path: []const u8) ![]u8 {
    return readFileFromDir(fs.cwd(), allocator, file_path);
}

pub fn readFileFromDir(dir: fs.Dir, allocator: mem.Allocator, file_path: []const u8) ![]u8 {
    return dir.readFileAlloc(allocator, file_path, std.math.maxInt(usize));
}

// Reads entire file into a single string with all newlines stripped
pub fn readFileFlat(allocator: mem.Allocator, file_path: []const u8) ![]u8 {
    return readFileFlatFromDir(fs.cwd(), allocator, file_path);
}

pub fn readFileFlatFromDir(dir: fs.Dir, allocator: mem.Allocator, file_path: []const u8) ![]u8 {
    const content = try dir.readFileAlloc(allocator, file_path, std.math.maxInt(usize));

    var write: usize = 0;
    for (content) |c| {
        if (c != '\n') {
            content[write] = c;
            write += 1;
        }
    }

    if (allocator.resize(content, write)) {
        return content[0..write];
    }
    const shrunk = try allocator.dupe(u8, content[0..write]);
    allocator.free(content);
    return shrunk;
}

/// Reads entire file as raw bytes
pub fn readFileBytes(allocator: mem.Allocator, file_path: []const u8) ![]u8 {
    return readFileBytesFromDir(fs.cwd(), allocator, file_path);
}

pub fn readFileBytesFromDir(dir: fs.Dir, allocator: mem.Allocator, file_path: []const u8) ![]u8 {
    return dir.readFileAlloc(allocator, file_path, std.math.maxInt(usize));
}

// Reads file into lines
// Zero-copy: all lines are slices into one backing buffer
// result.deinit() when done
pub fn readLines(allocator: mem.Allocator, file_path: []const u8) !Lines {
    return readLinesFromDir(fs.cwd(), allocator, file_path);
}

pub fn readLinesFromDir(dir: fs.Dir, allocator: mem.Allocator, file_path: []const u8) !Lines {
    const content = try dir.readFileAlloc(allocator, file_path, std.math.maxInt(usize));
    errdefer allocator.free(content);

    var line_list = std.ArrayList([]const u8).init(allocator);
    errdefer line_list.deinit();

    var iter = mem.splitScalar(u8, content, '\n');
    while (iter.next()) |line| {
        const trimmed = if (line.len > 0 and line[line.len - 1] == '\r')
            line[0 .. line.len - 1]
        else
            line;
        try line_list.append(trimmed);
    }

    // splitScalar produces a trailing empty element when content ends with \n
    if (line_list.items.len > 0 and line_list.items[line_list.items.len - 1].len == 0) {
        _ = line_list.pop();
    }

    return .{
        .items = try line_list.toOwnedSlice(),
        ._backing = content,
        .allocator = allocator,
    };
}

// Reads file line-by-line via streaming
// Heap-allocates the read buffer
// Good for large files never holds full content in memory
pub fn readLinesStreaming(allocator: mem.Allocator, file_path: []const u8, context: anytype, processLine: fn (@TypeOf(context), usize, []const u8) anyerror!void) !void {
    return readLinesStreamingFromDir(allocator, fs.cwd(), file_path, context, processLine);
}

pub fn readLinesStreamingFromDir(allocator: mem.Allocator, dir: fs.Dir, file_path: []const u8, context: anytype, processLine: fn (@TypeOf(context), usize, []const u8) anyerror!void) !void {
    const file = try dir.openFile(file_path, .{});
    defer file.close();

    const read_buffer = try allocator.alloc(u8, 1024 * 1024);
    defer allocator.free(read_buffer);

    var file_reader = file.reader(read_buffer);
    const reader = &file_reader.interface;

    var line_num: usize = 0;
    while (true) {
        const line = (try reader.takeDelimiter('\n')) orelse break;
        const trimmed = if (line.len > 0 and line[line.len - 1] == '\r')
            line[0 .. line.len - 1]
        else
            line;
        try processLine(context, line_num, trimmed);
        line_num += 1;
    }
}

// Reads non-empty lines from a file
pub fn readNonEmptyLines(allocator: mem.Allocator, file_path: []const u8) !Lines {
    return readNonEmptyLinesFromDir(fs.cwd(), allocator, file_path);
}

pub fn readNonEmptyLinesFromDir(dir: fs.Dir, allocator: mem.Allocator, file_path: []const u8) !Lines {
    const content = try dir.readFileAlloc(allocator, file_path, std.math.maxInt(usize));
    errdefer allocator.free(content);

    var line_list = std.ArrayList([]const u8).init(allocator);
    errdefer line_list.deinit();

    var iter = mem.splitScalar(u8, content, '\n');
    while (iter.next()) |line| {
        const trimmed_cr = if (line.len > 0 and line[line.len - 1] == '\r')
            line[0 .. line.len - 1]
        else
            line;

        if (mem.trim(u8, trimmed_cr, " \t\r").len > 0) {
            try line_list.append(trimmed_cr);
        }
    }

    return .{
        .items = try line_list.toOwnedSlice(),
        ._backing = content,
        .allocator = allocator,
    };
}

// Counts lines by scanning for '\n'
pub fn countLines(allocator: mem.Allocator, file_path: []const u8) !usize {
    return countLinesInDir(fs.cwd(), allocator, file_path);
}

pub fn countLinesInDir(dir: fs.Dir, allocator: mem.Allocator, file_path: []const u8) !usize {
    const content = try dir.readFileAlloc(allocator, file_path, std.math.maxInt(usize));
    defer allocator.free(content);

    if (content.len == 0) return 0;

    var count: usize = 0;
    for (content) |c| {
        if (c == '\n') count += 1;
    }

    // if file doesn't end with \n, the last line still counts
    if (content[content.len - 1] != '\n') count += 1;

    return count;
}
