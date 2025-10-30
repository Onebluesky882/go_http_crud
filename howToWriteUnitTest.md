Arrange → Act → Assert

Arrange: เตรียมค่า input หรือ mock object

Act: เรียกฟังก์ชันที่ต้องการ test

Assert: ตรวจสอบผลลัพธ์ (expect vs actual)

```go

package mypackage

import "testing"

func Add(a, b int) int {
    return a + b
}

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("expected 5, got %d", result)
    }
}
```
