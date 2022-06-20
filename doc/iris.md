[![���˵���Ҳ����](https://iris-go.com/images/blacklivesmatter_banner.png)](https://support.eji.org/give/153413/#!/donation/checkout)

> ����һ��**�����еİ汾**�������ע���������İ汾 [v12.2.0](HISTORY.md#Next)�������ʹ���ȶ��汾����鿴 [v12.1.8 ��֧](https://github.com/kataras/iris/tree/v12.1.8) ��
>
> ![](https://iris-go.com/images/cli.png) �������Թٷ���[Iris�����й���](https://github.com/kataras/iris-cli)��

# Iris Web Framework

Iris �ǻ��� Go ��д��һ�����٣��򵥵�������ȫ�ҷǳ���Ч�� Web ��ܡ� 

��Ϊ������һ����վ�� API �ṩ��һ���ǳ����б�����������ʹ�õĻ�����

���� [������������� Iris](https://iris-go.com/testimonials/)��ͬʱ��ӭ��λΪ�˿�Դ��Ŀ���� **[star](https://github.com/kataras/iris/stargazers)**��

[![](https://iris-go.com/images/reviews.gif)](https://iris-go.com/testimonials/)

[![Benchmarks: Jul 18, 2020 at 10:46am (UTC)](https://iris-go.com/images/benchmarks.svg)](https://github.com/kataras/server-benchmarks)

## ?? ��ʼѧϰ Iris

```sh
# ��װIris��https://github.com/kataras/iris/wiki/Installation
$ go get github.com/kataras/iris/v12@master
# ����main.go�ļ����Ѵ������´���
$ cat main.go
```

```go
package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()

	booksAPI := app.Party("/books")
	{
		booksAPI.Use(iris.Compression)

		// GET: http://localhost:8080/books
		booksAPI.Get("/", list)
		// POST: http://localhost:8080/books
		booksAPI.Post("/", create)
	}

	app.Listen(":8080")
}

// Book example.
type Book struct {
	Title string `json:"title"`
}

func list(ctx iris.Context) {
	books := []Book{
		{"Mastering Concurrency in Go"},
		{"Go Design Patterns"},
		{"Black Hat Go"},
	}

	ctx.JSON(books)
	// ��ʾ: �ڷ��������ȼ��Ϳͻ��������н�����ӦЭ�̣�
	// �Դ������� ctx.JSON:
	// ctx.Negotiation().JSON().MsgPack().Protobuf()
	// ctx.Negotiate(books)
}

func create(ctx iris.Context) {
	var b Book
	err := ctx.ReadJSON(&b)
	// ��ʾ: ʹ�� ctx.ReadBody(&b) ���棬�����������͵����
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("Book creation failure").DetailErr(err))
		// ��ʾ: ������д��ı���plain text��������Ӧ��
        // ��ʹ�� ctx.StopWithError(code, err) 
		return
	}

	println("Received Book: " + b.Title)

	ctx.StatusCode(iris.StatusCreated)
}
```

ͬ���أ���**MVC**�� :

```go
import "github.com/kataras/iris/v12/mvc"
```

```go
m := mvc.New(booksAPI)
m.Handle(new(BookController))
```

```go
type BookController struct {
	/* dependencies */
}

// GET: http://localhost:8080/books
func (c *BookController) Get() []Book {
	return []Book{
		{"Mastering Concurrency in Go"},
		{"Go Design Patterns"},
		{"Black Hat Go"},
	}
}

// POST: http://localhost:8080/books
func (c *BookController) Post(b Book) int {
	println("Received Book: " + b.Title)

	return iris.StatusCreated
}
```

**����** ���� Iris web ����:

```sh
$ go run main.go
> Now listening on: http://localhost:8080
> Application started. Press CTRL+C to shut down.
```

Books **�б��ѯ** :

```sh
$ curl --header 'Accept-Encoding:gzip' http://localhost:8080/books

[
  {
    "title": "Mastering Concurrency in Go"
  },
  {
    "title": "Go Design Patterns"
  },
  {
    "title": "Black Hat Go"
  }
]
```

**����** �µ�Book:

```sh
$ curl -i -X POST \
--header 'Content-Encoding:gzip' \
--header 'Content-Type:application/json' \
--data "{\"title\":\"Writing An Interpreter In Go\"}" \
http://localhost:8080/books

> HTTP/1.1 201 Created
```

����**����**��Ӧ��չʾ�����ӣ�

```sh
$ curl -X POST --data "{\"title\" \"not valid one\"}" \
http://localhost:8080/books

> HTTP/1.1 400 Bad Request

{
  "status": 400,
  "title": "Book creation failure"
  "detail": "invalid character '\"' after object key",
}
```

</details>

[![run in the browser](https://img.shields.io/badge/Run-in%20the%20Browser-348798.svg?style=for-the-badge&logo=repl.it)](https://bit.ly/2YJeSZe)

Iris ���������꾡�� **[ʹ���ĵ�](https://github.com/kataras/iris/wiki)** �������������ɵ�ʹ�ô˿�ܡ�

<!-- ![](https://media.giphy.com/media/Ur8iqy9FQfmPuyQpgy/giphy.gif) -->

Ҫ�˽����ϸ�ļ����ĵ�����������ǵ� [godocs](https://pkg.go.dev/github.com/kataras/iris/v12@master)�������ҪѰ�Ҵ���ʾ���������Ե��ֿ�� [./_examples](_examples) ��Ŀ¼�»�ȡ��

### ��ϲ��������ʱ�Ķ���

<a href="https://iris-go.com/#book"> <img alt="Book cover" src="https://iris-go.com/images/iris-book-cover-sm.jpg?v=12" /> </a>

[![follow Iris web framework on twitter](https://img.shields.io/twitter/follow/iris_framework?color=ee7506&logoColor=ee7506&style=for-the-badge)](https://twitter.com/intent/follow?screen_name=iris_framework)

������[��ȡ](https://www.iris-go.com/#ebookDonateForm)PDF�汾�����߷���**����ͼ��**�������뵽Iris�Ŀ����С�

## ?? ����

���ǻ�ӭ��ΪIris����������ף���Ҫ֪�����ΪIris��Ŀ�����ף���鿴[CONTRIBUTING.md](CONTRIBUTING.md)��

[����������](https://github.com/kataras/iris/graphs/contributors)

## ?? ��ȫ©��

����������� Iris ���ڰ�ȫ©�����뷢�͵����ʼ��� [iris-go@outlook.com](mailto:iris-go@outlook.com)�����а�ȫ©������õ���ʱ�����

## ?? ��ԴЭ�飨License��

����Go���Ե�Э��һ��������ĿҲ���� [BSD 3-clause license](LICENSE)��

��Ŀ���� "Iris" �����������ϣ���񻰡�

<!-- ## Stargazers over time

[![Stargazers over time](https://starchart.cc/kataras/iris.svg)](https://starchart.cc/kataras/iris) -->