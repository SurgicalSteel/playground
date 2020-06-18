const EventEmitter = require('events');
const emitter = new EventEmitter();

emitter.on('divideByZero',(arg) => {
    console.log(`Invalid argument, divide by zero exception. a=${arg.a} , b=${arg.b}`);
});

function Divide(a,b){
    if(b == 0){
        emitter.emit('divideByZero',{a:a,b:b});
    }
    else{
        console.log(a/b);
    }
}

exports.Divide = Divide;