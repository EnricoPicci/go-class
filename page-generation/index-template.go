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
      width: 90%%;
      height: auto;
    }
  </style>
</head>
<body>
  <h1>Go-class</h1>
  <h2>A journey to learn Go programming language</h2>
  <p style="display: inline">Created by Enrico Piccinin</p>
  <a
    href="mailto:enrico.piccinin@gmail.com?subject = Go-class"
    style="margin-left: 20"
    >enrico.piccinin@gmail.com</a
  >
  <img src="./assets/img/go-class.png" class="responsive" />
  <ul>
    %v
  </ul>
  <p>
    This work is licensed under the Creative Commons Attribution-NonCommercial-ShareAlike (CC BY-NC-SA) license. 
    See <a href="./LICENSE.txt">license.txt</a> for details.
  </p>
</body>
`
