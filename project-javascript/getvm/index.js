// const obj = { session_id: true };

// Function to get details and handle form submission
function getdetails() {
  // Fetch input values
  const hardware = document.getElementById("hardware").value;
  const version = document.getElementById("version").value;
  const count = document.getElementById("count").value;
  const caseno = document.getElementById("caseno").value;

  const formData = {
    hardwarevalue: hardware,
    version: version,
    count: count,
    caseNo: caseno,
  };
  //   console.log("Form Data:", formData); // Optional: Log form data for debugging
  let session_id = true;
  sessionStorage.setItem("session_id", session_id);
  sessionStorage.setItem("Form Data", JSON.stringify(formData)); // Optional: Log form data for debugging

  // Disable the button and provide feedback
  const button = document.getElementById("button");
  button.disabled = true;
  button.textContent = "Loading...";
  button.style.color = "black";
  button.style.backgroundColor = "#4efa3a";

  // Style the page background and text
  //   const body = document.body;
  //   body.style.backgroundColor = "#333";
  //   body.style.color = "white";

  // Display a popup
  const popup = document.getElementById("popup");
  popup.style.display = "block";

  // Close the popup
  const closePopupButton = document.getElementById("close-popup");
  closePopupButton.addEventListener("click", () => {
    popup.style.display = "none";
    // location.reload();
  });

  //   // Close the popup when clicking outside the popup content
  //   popup.addEventListener("click", (event) => {
  //     if (event.target === popup) {
  //       popup.style.display = "none";
  //     }
  //   });

  para = document.querySelector("p");

  const session_id_1 = sessionStorage.getItem("session_id");
  if (session_id_1 === "true") {
    para.textContent = `Please wait while your info${hardware} appliance is being created`;
  } else {
    para.textContent = `Enter Values`;
    console.log("Enter values");
  }
}

// OUTSIDE FUntion

const button = document.getElementById("button");
const body = document.getElementById("body");
let isDarkMode = false; // State to track dark mode

button.addEventListener("click", () => {
  isDarkMode = true; // Enable dark mode
  body.style.backgroundColor = "#333";
  body.style.color = "white";
});

button.addEventListener("mouseover", (event) => {
  if (!isDarkMode) {
    body.style.backgroundColor = "#333";
    body.style.color = "white";
  }
});

button.addEventListener("mouseout", (event) => {
  if (!isDarkMode) {
    body.style.backgroundColor = "white";
    body.style.color = "black";
  }
});
