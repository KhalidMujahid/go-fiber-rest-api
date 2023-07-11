const fname = document.querySelector("#fname");
const lname = document.querySelector("#lname");
const age = document.querySelector("#age");
const button = document.querySelector("button");

async function getData() {
  await fetch("/users")
    .then((res) => res.json())
    .then((data) => console.log(data))
    .catch((error) => console.error(error));
}

getData();

button.addEventListener("click", (e) => {
  e.preventDefault();

  const value = { fname: fname.value, lname: lname.value, age: age.value };

  sendToBackEnd(value);
});

async function sendToBackEnd(value) {
  await fetch("/users", {
    method: "POST",
    headers: {
      "Content-type": "application/json",
    },
    body: JSON.stringify(value),
  })
    .then((res) => res.json())
    .then((data) => console.log(data))
    .catch((error) => console.error(error));

  fname.value = "";
  lname.value = "";
  age.value = "";
}
