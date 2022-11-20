# jisx4061

Package jisx4061 implements [JIS X 4061](https://ja.wikipedia.org/wiki/%E6%97%A5%E6%9C%AC%E8%AA%9E%E6%96%87%E5%AD%97%E5%88%97%E7%85%A7%E5%90%88%E9%A0%86%E7%95%AA)
Japanese character string collation order.
It is commonly referred to as "辞書順(the dictionary order)", "50音順(the syllabic order), "あいうえお順(the a-i-u-e-o order)".

## SYNOPSIS

```go
	list := []string{
		"さどう",
		"さとうや",
		"サトー",
		"さと",
		"さど",
		"さとう",
		"さとおや",
	}
	jisx4061.Sort(list)
	for _, s := range list {
		fmt.Println(s)
	}
	// Output:
	// さと
	// さど
	// さとう
	// さどう
	// さとうや
	// サトー
	// さとおや
```

## REFERENCES

- [[JIS X 4061](https://ja.wikipedia.org/wiki/%E6%97%A5%E6%9C%AC%E8%AA%9E%E6%96%87%E5%AD%97%E5%88%97%E7%85%A7%E5%90%88%E9%A0%86%E7%95%AA)
