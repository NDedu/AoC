#ifndef C_UTILS_H
#define C_UTILS_H

#include <errno.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// Files

static inline char *cutils_read_all(const char *path, const char *caller, size_t *length) {

	if (length) {
		*length = 0;
	}

	FILE *file = fopen(path, "rb");
	if (!file) {
		fprintf(stderr, "%s: cannot open '%s': %s\n", caller, path, strerror(errno));
		return NULL;
	}

	size_t capacity = 0;
	if (fseek(file, 0, SEEK_END) == 0) {

		long size = ftell(file);
		if (size > 0) {
			capacity = (size_t)size + 2;
		}

		rewind(file);
	}

	if (capacity == 0) {
		capacity = 1 << 16;
	}

	size_t used = 0;
	char *content = (char *)malloc(capacity);
	if (!content) {
		fprintf(stderr, "%s: out of memory reading '%s'\n", caller, path);
		fclose(file);
		return NULL;
	}

	for (;;) {

		size_t room = capacity - used - 1;

		if (room == 0) {

			size_t grown = capacity * 2;
			char *bigger = (char *)realloc(content, grown);
			if (!bigger) {
				fprintf(stderr, "%s: out of memory reading '%s'\n", caller, path);
				free(content);
				fclose(file);
				return NULL;
			}

			content = bigger;
			capacity = grown;
			room = capacity - used - 1;
		}

		size_t read_bytes = fread(content + used, 1, room, file);
		used += read_bytes;

		if (read_bytes < room) {
			break;
		}
	}

	bool failed = ferror(file) != 0;
	fclose(file);

	if (failed) {
		fprintf(stderr, "%s: error reading '%s': %s\n", caller, path, strerror(errno));
		free(content);
		return NULL;
	}

	content[used] = '\0';

	if (used > 0 && content[used - 1] == '\n') {

		used--;
		if (used > 0 && content[used - 1] == '\r') {
			used--;
		}

		content[used] = '\0';
	}

	if (capacity - used > 4096) {

		char *tight = (char *)realloc(content, used + 1);
		if (tight) {
			content = tight;
		}
	}

	if (length) {
		*length = used;
	}

	return content;
}

static inline char *read_file(const char *path) {

	return cutils_read_all(path, "read_file", NULL);
}

static inline char **read_lines(const char *path, size_t *count) {

	if (count) {
		*count = 0;
	}

	size_t length = 0;
	char *content = cutils_read_all(path, "read_lines", &length);
	if (!content) {
		return NULL;
	}

	size_t rows = 0;
	if (length > 0) {

		rows = 1;
		for (size_t i = 0; i < length; i++) {

			if (content[i] == '\n') {
				rows++;
			}
		}
	}

	char **lines = (char **)malloc((rows + 1) * sizeof(*lines) + length + 1);
	if (!lines) {
		fprintf(stderr, "read_lines: out of memory reading '%s'\n", path);
		free(content);
		return NULL;
	}

	char *text = (char *)(lines + rows + 1);
	memcpy(text, content, length + 1);
	free(content);

	size_t written = 0;
	size_t start = 0;

	for (size_t i = 0; i <= length && written < rows; i++) {

		if (i != length && text[i] != '\n') {
			continue;
		}

		size_t span = i - start;
		if (span > 0 && text[start + span - 1] == '\r') {
			span--;
		}

		text[start + span] = '\0';
		lines[written++] = text + start;
		start = i + 1;
	}

	lines[written] = NULL;

	if (count) {
		*count = written;
	}

	return lines;
}

static inline void free_lines(char **lines) {

	free(lines);
}

#endif // C_UTILS_H
