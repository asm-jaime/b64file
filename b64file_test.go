package b64file

import (
	"errors"
	"reflect"
	"testing"
)

const dataCorrect = `data:image/jpeg;base64,/9j/4AAQSkZJRgABAQEAkACQAAD/2wBD
AAgGBgcGBQgHBwcJCQgKDBQNDAsLDBkSEw8UHRofHh0aHBwgJC4nICIsIxwcKDcpLDAxNDQ0Hyc5
PTgyPC4zNDL/2wBDAQkJCQwLDBgNDRgyIRwhMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIy
MjIyMjIyMjIyMjIyMjIyMjIyMjL/wAARCAABAAEDASIAAhEBAxEB/8QAHwAAAQUBAQEBAQEAAAAA
AAAAAAECAwQFBgcICQoL/8QAtRAAAgEDAwIEAwUFBAQAAAF9AQIDAAQRBRIhMUEGE1FhByJxFDKB
kaEII0KxwRVS0fAkM2JyggkKFhcYGRolJicoKSo0NTY3ODk6Q0RFRkdISUpTVFVWV1hZWmNkZWZn
aGlqc3R1dnd4eXqDhIWGh4iJipKTlJWWl5iZmqKjpKWmp6ipqrKztLW2t7i5usLDxMXGx8jJytLT
1NXW19jZ2uHi4+Tl5ufo6erx8vP09fb3+Pn6/8QAHwEAAwEBAQEBAQEBAQAAAAAAAAECAwQFBgcI
CQoL/8QAtREAAgECBAQDBAcFBAQAAQJ3AAECAxEEBSExBhJBUQdhcRMiMoEIFEKRobHBCSMzUvAV
YnLRChYkNOEl8RcYGRomJygpKjU2Nzg5OkNERUZHSElKU1RVVldYWVpjZGVmZ2hpanN0dXZ3eHl6
goOEhYaHiImKkpOUlZaXmJmaoqOkpaanqKmqsrO0tba3uLm6wsPExcbHyMnK0tPU1dbX2Nna4uPk
5ebn6Onq8vP09fb3+Pn6/9oADAMBAAIRAxEAPwD5/ooooA//2Q==`
const dataIncorrect = `data:image/png;base64,dfdfdf`
const dataIncorrectFormat = `data:image/png;base64,/9j/4AAQSkZJRgABAQEAkACQA
AD/2wBDAAgGBgcGBQgHBwcJCQgKDBQNDAsLDBkSEw8UHRofHh0aHBwgJC4nICIsIxwcKDcpLDAxN
DQ0Hyc5PTgyPC4zNDL/2wBDAQkJCQwLDBgNDRgyIRwhMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyM
jIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjL/wAARCAABAAEDASIAAhEBAxEB/8QAHwAAAQUBAQEBA
QEAAAAAAAAAAAECAwQFBgcICQoL/8QAtRAAAgEDAwIEAwUFBAQAAAF9AQIDAAQRBRIhMUEGE1FhB
yJxFDKBkaEII0KxwRVS0fAkM2JyggkKFhcYGRolJicoKSo0NTY3ODk6Q0RFRkdISUpTVFVWV1hZW
mNkZWZnaGlqc3R1dnd4eXqDhIWGh4iJipKTlJWWl5iZmqKjpKWmp6ipqrKztLW2t7i5usLDxMXGx
8jJytLT1NXW19jZ2uHi4+Tl5ufo6erx8vP09fb3+Pn6/8QAHwEAAwEBAQEBAQEBAQAAAAAAAAECA
wQFBgcICQoL/8QAtREAAgECBAQDBAcFBAQAAQJ3AAECAxEEBSExBhJBUQdhcRMiMoEIFEKRobHBC
SMzUvAVYnLRChYkNOEl8RcYGRomJygpKjU2Nzg5OkNERUZHSElKU1RVVldYWVpjZGVmZ2hpanN0d
XZ3eHl6goOEhYaHiImKkpOUlZaXmJmaoqOkpaanqKmqsrO0tba3uLm6wsPExcbHyMnK0tPU1dbX2
Nna4uPk5ebn6Onq8vP09fb3+Pn6/9oADAMBAAIRAxEAPwD5/ooooA//2Q==`
const dataUnsupportedFormat = `data:image/bmp;base64,Qk06AAAAAAAAADYAAAAoAAA
AAQAAAAEAAAABABgAAAAAAAQAAAAlFgAAJRYAAAAAAAAAAAAAAAAAAA==`

func TestB64ToFile(t *testing.T) {
	tests := []struct {
		name string
		path string
		data string
		want error
	}{
		{"correct data",
			"./data.correct",
			dataCorrect,
			nil,
		},
		{"incorrect format in b64 prefix",
			"./data.incorrect_prefix_format",
			dataIncorrectFormat,
			nil,
		},
		{"unsupported format bmp",
			"./data.unsupported_format",
			dataUnsupportedFormat,
			errors.New("image: unknown format"),
		},
		{"incorrect data",
			"./data.incorrect",
			dataIncorrect,
			errors.New("unexpected EOF"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := B64ToFile(tt.path, tt.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("B64ToFile = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileToB64(t *testing.T) {
	tests := []struct {
		name string
		path string
		want error
	}{
		{"file correct data",
			"./data.correct.jpeg",
			nil,
		},
		{"file incorrect format in b64 prefix",
			"./data.incorrect_prefix_format.jpeg",
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got := FileToB64(tt.path)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("B64ToFile = %v, want %v", got, tt.want)
			}
		})
	}
}
