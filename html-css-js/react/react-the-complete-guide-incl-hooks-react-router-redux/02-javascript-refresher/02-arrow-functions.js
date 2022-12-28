// function myFunc() {
//   ...;
// }

// const myFunc = () => {
//   ...;
// }

function printMyName(name) {
  console.log(name);
}

printMyName('Max');

// const constPrintMyName = (name) => {
const constPrintMyName = name => {
  console.log(name);
};

constPrintMyName('Max');

const multiply = (number) => number * 2;

console.log(multiply(2));