package main

import (
	"log"
	"strconv"
	"strings"
	"testing"
)

func concatPlusEquals(values []string) string {
	s := ""
	for _, value := range values {
		s += value
	}
	return s
}

func concatJoin(values []string) string {
	return strings.Join(values, "")
}

func concatStringsBuilderV1(values []string) string {
	sb := strings.Builder{}
	for _, value := range values {
		_, err := sb.WriteString(value)
		if err != nil {
			log.Fatal(err)
		}
	}
	return sb.String()
}

func concatStringsBuilderV2(values []string) string {
	total := 0
	for i := 0; i < len(values); i++ {
		total += len(values[i])
	}

	sb := strings.Builder{}
	sb.Grow(total)
	for _, value := range values {
		_, _ = sb.WriteString(value)
	}
	return sb.String()
}

func createInitialStringList() []string {
	var bl []string
	for i := 0; i < 10_000; i++ {
		bl = append(bl, strconv.Itoa(i))
	}

	return bl
}

func TestStrings(t *testing.T) {
	stringList := createInitialStringList()

	result1 := concatPlusEquals(stringList)
	result2 := concatJoin(stringList)
	result3 := concatStringsBuilderV1(stringList)
	result4 := concatStringsBuilderV2(stringList)

	if result1 != result2 || result1 != result3 || result1 != result4 {
		t.Error("expected all results to be equal")
	}

}

func BenchmarkStrings(b *testing.B) {
	stringList := createInitialStringList()

	b.Run("concatPlusEquals", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = concatPlusEquals(stringList)
		}
		b.ReportAllocs()
	})

	b.Run("concatJoin", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = concatJoin(stringList)
		}
		b.ReportAllocs()
	})

	b.Run("concatStringBuilderV1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = concatStringsBuilderV1(stringList)
		}
		b.ReportAllocs()
	})

	b.Run("concatStringBuilderV2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = concatStringsBuilderV2(stringList)
		}
		b.ReportAllocs()
	})

}
