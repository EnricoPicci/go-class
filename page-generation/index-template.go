package pagegeneration

var indexTemplate = `
<head>
  <title>Go-class</title>
  <style>
    ul {
      -webkit-column-count: 2;
      -moz-column-count: 2;
      column-count: 2;
    }
  </style>
  <style>
    .responsive {
      width: 100%;
      height: auto;
    }
  </style>
</head>
<body>
  <h1>Go-class</h1>
  <h2>A journey to learn Go programming language</h2>
  <p>Created by Enrico Piccinin</p>
  <a href="mailto:enrico.piccinin@gmail.com?subject = Go-class"
    >enrico.piccinin@gmail.com</a
  >
  <img src="./assets/img/go-class.png" class="responsive" />
  <ul>
    %v
  </ul>
</body>
`
