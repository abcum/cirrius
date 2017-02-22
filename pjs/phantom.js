// Specify the encoding
phantom.outputEncoding = "UTF-8";

// Initialise webserver
var server = require('webserver').create();

// Listen for connections
server.listen('localhost:8080', function(request, response) {

    if (!request.headers.Url) {
        response.statusCode = 500;
        response.close();
    }

    var page = new WebPage();

    // Specify page settings
    page.settings.userAgent = "Cirrius";
    page.settings.loadImages = false;
    page.settings.resourceTimeout = 3000;

    page.open(request.headers.Url, function(status) {

        if (status === 'success') {

            page.evaluate(function() {
                Array.prototype.slice.call( document.getElementsByTagName("script") ).forEach(function(s) {
                    s.parentNode.removeChild(s);
                });
            });

            response.headers = {
                'Content-Length': request.content.length,
                'Content-Type': 'text/html;',
                'Connection': 'keep-alive',
                'Keep-Alive': 'timeout=60, max=300',
            };

            response.statusCode = 200;
            response.write(page.content);
            response.close();

        }

        if (status !== 'success') {
            response.statusCode = 500;
            response.close();
        }

        page.close();

    });

});
