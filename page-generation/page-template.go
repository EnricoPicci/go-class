package pagegeneration

var pageTemplate = `
<head>
  <title>%v</title>
</head>
<body>
  <h1>%v</h1>

  <iframe
    src="https://docs.google.com/presentation/d/%v/embed?start=false&loop=false&delayms=3000"
    frameborder="0"
    width="960"
    height="569"
    allowfullscreen="true"
    mozallowfullscreen="true"
    webkitallowfullscreen="true"
  ></iframe>
  <form action="../index.html" style="display: inline">
    <input type="submit" value="Go to class modules index" />
  </form>
  <p>
    Go-class created by Enrico Piccinin (<a
      href="mailto:enrico.piccinin@gmail.com?subject = Go-class"
      >enrico.piccinin@gmail.com</a
    >)
  </p>
  <p>
    This work is licensed under the Creative Commons Attribution-NonCommercial-ShareAlike (CC BY-NC-SA) license. 
    See <a href="../LICENSE.txt">license.txt</a> for details.
  </p>
</body>
`
