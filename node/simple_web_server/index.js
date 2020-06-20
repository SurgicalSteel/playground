const http = require('http');

// http.Server inherits from net.Server which is an eventEmitter
const server = http.createServer((request,response)=>{

    //the router here is full of mess, to get organized, use expressJS as a node framework.
    if(request.url==='/' && request.method==='GET'){
        response.write('Hello world!');
        response.end();
        return;
    }

    if(request.url === '/members' && request.method==='GET'){
        response.write(
            JSON.stringify(
                [
                    {
                        name: 'Angela Gossow',
                        position : 'Vocal'
                    },
                    {
                        name: 'Michael Amott',
                        position : 'Guitar'
                    },
                    {
                        name: 'Christopher Amott',
                        position : 'Guitar'
                    },
                    {
                        name: 'Sharlee D\'Angelo',
                        position : 'Bass'
                    },
                    {
                        name: 'Daniel Erlandsson',
                        position : 'Drum'
                    }
                ]
            )
        );
        response.end();
        return;
    }

    response.writeHead(404,{'Content-Type':'text/html'});

    response.write('<h1>404</h1></br><h1>Oops, the page you requested was not found. Are you lost?</h1>');
    response.end();
});   

//handles on each new incoming connection, 'connection' is the eventEmitter with 1 arg (socket)
server.on('connection',(socket)=>{
    console.log('A new connection');
    //console.log(socket.remoteAddress);
    //console.log(socket.remotePort)
});


server.listen(9999);
console.log('Listening on port 9999')