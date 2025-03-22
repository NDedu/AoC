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

pub fn appendToFile(file_path: []const u8, content: []const u8) !void {
    // Open the file in append mode, create if it doesn't exist
    const file = try std.fs.cwd().createFile(file_path, .{ .truncate = false, .read = false, .exclusive = false });
    defer file.close();

    // Seek to the end of the file
    try file.seekFromEnd(0);

    // Check if the last character is a newline, if not, add one
    const stat = try file.stat();
    if (stat.size > 0) {
        var buffer: [1]u8 = undefined;
        try file.seekTo(stat.size - 1);
        _ = try file.readAll(&buffer);
        try file.seekFromEnd(0);

        if (buffer[0] != '\n') {
            try file.writeAll("\n");
        }
    }

    // Write the content
    try file.writeAll(content);

    // Ensure the file ends with a newline
    if (content.len > 0 and content[content.len - 1] != '\n') {
        try file.writeAll("\n");
    }
}

pub fn writeToLine(allocator: std.mem.Allocator, file_path: []const u8, line_number: usize, content: []const u8) !void {
    // Read all lines from the file
    var lines = try readLines(allocator, file_path);
    defer {
        for (lines.items) |line| {
            allocator.free(line);
        }
        lines.deinit();
    }

    // Create a new file writer
    const file = try std.fs.cwd().createFile(file_path, .{});
    defer file.close();

    // Write all lines, replacing the specified line
    for (lines.items, 0..) |line, i| {
        if (i == line_number) {
            // Write the new content instead of the original line
            try file.writeAll(content);
            if (content.len > 0 and content[content.len - 1] != '\n') {
                try file.writeAll("\n");
            }
        } else {
            // Write the original line
            try file.writeAll(line);
            try file.writeAll("\n");
        }
    }

    // If the line number is beyond the end of the file, append it
    if (line_number >= lines.items.len) {
        // Add newlines to reach the desired line if needed
        if (lines.items.len > 0) {
            const lines_to_add = line_number - lines.items.len;
            var i: usize = 0;
            while (i < lines_to_add) : (i += 1) {
                try file.writeAll("\n");
            }
        }

        // Write the content
        try file.writeAll(content);
        if (content.len > 0 and content[content.len - 1] != '\n') {
            try file.writeAll("\n");
        }
    }
}
