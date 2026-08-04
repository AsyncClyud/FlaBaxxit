"use strict";

async function SendRegisterRequest() {
  const is_agree = document.getElementById("agree_button").value
  if (is_agree != "agree") {
    document.getElementById("status").textContent = "You need to agree with Privacy Policy and User Agreement!"
    return
  }

  const username = document.getElementById("username").value
  const password = document.getElementById("password").value
  const turnstile_token = turnstile.getResponse()

  const response = await fetch("/auth/register", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      Username: username,
      Password: password,
      Turnstile_token: turnstile_token
    })
  })
  if (response.ok) {
    const message = await response.json()
    document.getElementById("status").textContent = message.Message
    await new Promise(r => setTimeout(r, 2000));
    window.location.replace("/")
    }
  else {
    const message = await response.json()
    document.getElementById("status").textContent = message.Message
    turnstile.reset()
  }
}
