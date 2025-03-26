const std = @import("std");
const util = @import("../../zigUtils");

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    const allocator = gpa.allocator();
    defer _ = gpa.deinit();

    const input = try util.readFile(allocator, "../input/day1.txt");
    defer allocator.free(input);
    std.debug.print("Line: {s}\n", .{input});
}
