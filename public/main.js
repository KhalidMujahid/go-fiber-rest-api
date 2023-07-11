const fname = document.querySelector("#fname");
const lname = document.querySelector("#lname");
const age = document.querySelector("#age");
const button = document.querySelector("button");
const users = document.querySelector(".users");

async function getData() {
  await fetch("/users")
    .then((res) => res.json())
    .then((data) => {
      data.forEach((d) => {
        users.innerHTML += `
        <div class="user">
          <h5>Name: ${d.fname} ${d.lname}</h5>
          <p>Age: ${d.age}</p>
        </div>
      `;
      });
    })
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
    .then(() => {
      alert("Data sent successfully...");
      window.location.reload();
    })
    .catch((error) => console.error(error));

  fname.value = "";
  lname.value = "";
  age.value = "";
}
