const EventEmitter = require('events');

class Person extends EventEmitter{
    constructor(name,age){
        super();
        this.name=name;
        this.age=age;
    }

    getName(){
        if(this.name === ''){
            this.emit('invalidName',{name:this.name});
        }
        return this.name;
    }

    getAge(){
        if(this.age < 0){
            this.emit('invalidAge',{age:this.age});
        }
        return this.age;
    }

}

module.exports=Person;