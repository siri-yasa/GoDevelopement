# creating the Html File
file_html = open("demo.html", "w")
# adding input data to the  html file 
file_html.write('''<html>
<head>
<title>HTML File</title>
</head> 
<body>
<h1>Welcome world</h1>           
<p>I have a interview call</p> 
</body>
</html>''')
# saving the data in to the html
file_html.close()

             