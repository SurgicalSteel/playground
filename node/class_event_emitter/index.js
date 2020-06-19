const Person = require('./person');

var michael = new Person('Michael Amott',-48);

michael.on('invalidAge',(arg)=>{
    console.log('Error! The person\'s age is invalid! '+JSON.stringify(arg));
});

console.log('Name : '+michael.getName());
console.log('Age : '+michael.getAge());


