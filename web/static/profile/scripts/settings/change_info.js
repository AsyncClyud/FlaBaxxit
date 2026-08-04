"use strict";

async function FetchUserInfo() {
  const response = await fetch("/api/auth", {
    method: "GET",
    headers: { Accept: "application/json" }
  })
  if (response.ok) {
    const data = await response.json()
    const id = data.userID
  const request = await fetch(`/api/profile/${id}`, {
    method: "GET",
    headers: { Accept: "application/json" },
  })
    if (request.ok) {
      try{
        const data = JSON.parse(await request.json())
        document.getElementById("username").textContent = `${data.Username}`
        document.getElementById("bio").textContent = `${data.Bio}`
      }
      catch(fetch_error) {
        document.getElementById("username").textContent = `Failed to fetch username`
        document.getElementById("bio").textContent = `Failed to fetch username`
      }
    }
  }
}

async function SendUpdateUsernameRequest() {
  const new_username = document.getElementById("input_username").value
  const response = await fetch("/api/profile/username", {
    method: "PUT",
    headers: { Accept: "application/json" },
    body: JSON.stringify({
      Username: new_username
    })

  })
  if (response.ok) {
    const message = await response.json()
    document.getElementById("message_username").textContent = message.Message
  }
  else {
    const message = await response.json()
    document.getElementById("message_username").textContent = message.Message
  }
}

async function SendUpdateBioRequest() {
  const new_bio = document.getElementById("input_bio").value
  const response = await fetch("/api/profile/bio", {
    method: "PUT",
    headers: { Accept: "application/json" },
    body: JSON.stringify({
      Bio: new_bio
    })

  })
  if (response.ok) {
    const message = await response.json()
    document.getElementById("message_bio").textContent = message.Message
  }
  else {
    const message = await response.json()
    document.getElementById("message_bio").textContent = message.Message
  }
}

async function SendUpdatePasswordRequest() {
  const old_password = document.getElementById("input_old_password").value
  const new_password = document.getElementById("input_new_password").value
  const response = await fetch("/api/profile/password", {
    method: "PUT",
    headers: { Accept: "application/json" },
    body: JSON.stringify({
      Old_Password: old_password,
      New_Password: new_password
    })

  })
  if (response.ok) {
    const message = await response.json()
    document.getElementById("message_password").textContent = message.Message
  }
  else {
    const message = await response.json()
    document.getElementById("message_password").textContent = message.Message
  }
}

async function SendDeleteAccountRequest() {
  const response = await fetch("/api/users", {
    method: "DELETE"
  })
  const logout = await fetch("/api/logout", {
    method: "POST",
    credentials: "include"
  })
  if (response.ok) {
    const message = await response.json()
    document.getElementById("message_delete").textContent = message.Message
    await new Promise(r => setTimeout(r, 2000));
    window.location.replace("/")
  }
  else {
    const message = await response.json()
    document.getElementById("message_delete").textContent = message.Message
  }
}
