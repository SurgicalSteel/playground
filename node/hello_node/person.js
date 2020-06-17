class Person{
  constructor(name, age){
    this.name=name;
    this.age=age;
  }

  Greet(){
    console.log(`Hello, my name is ${this.name}, and I'm ${this.age}`);
  }
}
module.exports = Person;
