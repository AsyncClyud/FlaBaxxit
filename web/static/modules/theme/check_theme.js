"use strict";

function checkTheme() {
  const saved = localStorage.getItem("theme");
  if (saved === "light" || !saved) {
    document.documentElement.classList.add("light");
  }
  if (saved === "dark") {
    document.documentElement.classList.add("dark")
  }
}

export default checkTheme()
