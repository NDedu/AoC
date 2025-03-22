const std = @import("std");

pub fn readLines(allocator: std.mem.Allocator, file_path: []const u8) !std.ArrayList([]u8) {
    // Open the file
    const file = try std.fs.cwd().openFile(file_path, .{});
    defer file.close();

    // Create a buffered reader
    const buf_reader = std.io.bufferedReader(file.reader());
    const in_stream = buf_reader.reader();

    // Initialize an array list to store lines
    var lines = std.ArrayList([]u8).init(allocator);
    errdefer {
        for (lines.items) |line| {
            allocator.free(line);
        }
        lines.deinit();
    }

    // Read lines until EOF
    var buf: [1024]u8 = undefined;
    while (try in_stream.readUntilDelimiterOrEof(&buf, '\n')) |line| {
        // Create a copy of the line on the heap
        const line_copy = try allocator.dupe(u8, line);
        try lines.append(line_copy);
    }

    return lines;
}

pub fn readFile(allocator: std.mem.Allocator, file_path: []const u8) ![]u8 {
    // Open the file
    const file = try std.fs.cwd().openFile(file_path, .{});
    defer file.close();

    // Get the file size
    const file_size = try file.getEndPos();

    // Allocate a buffer to hold the entire file
    const buffer = try allocator.alloc(u8, file_size);
    errdefer allocator.free(buffer);

    // Read the entire file into the buffer
    const bytes_read = try file.readAll(buffer);

    // Make sure we read the entire file
    if (bytes_read != file_size) {
        return error.IncompleteRead;
    }

    return buffer;
}
