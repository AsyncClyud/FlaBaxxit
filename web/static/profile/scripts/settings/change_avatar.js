"use strict";

function SelectAvatar(avatar_id) {
  const picture = document.getElementById(`picture_${avatar_id}`)
  picture.setAttribute("class", "h-fit w-fit outline-[5px] outline-[#dee729]")
  picture.setAttribute("onclick", `UnselectAvatar(${avatar_id})`)
  const send_button = document.getElementById("change_avatar_button")
  send_button.setAttribute("onclick", `SendChangeAvatarRequest(${avatar_id})`)
}

function UnselectAvatar(avatar_id) {
  const picture = document.getElementById(`picture_${avatar_id}`)
  picture.setAttribute("class", "h-fit w-fit")
  picture.setAttribute("onclick", `SelectAvatar(${avatar_id})`)
  const send_button = document.getElementById("change_avatar_button")
  send_button.removeAttribute("onclick")
}

async function ShowChangeAvatarMenu() {
  const main_element = document.getElementById("actions")
  main_element.innerHTML = `
      <h2 class="h-[min-content] w-[75vw] bg-[#333c46] text-[gainsboro] text-[4vh] rounded-[10px] ml-auto mr-auto outline-[3px] outline-solid outline-[#151b23]">Account info</h2>
      <p class="m-[5px]">Select a new avatar:</p>
      <div class="h-[min-content] w-fit flex flex-row gap-4 mx-auto">
      <button class="h-fit w-fit" type="button" onclick="SelectAvatar(2)" id="picture_2"><img src="/static/images/2.webp" alt="Profile picture with ID 2" width="200px" height="200px"></button>
      <button class="h-fit w-fit" type="button" onclick="SelectAvatar(3)" id="picture_3"><img src="/static/images/3.webp" alt="Profile picture with ID 3" width="200px" height="200px"></button>
      <button class="h-fit w-fit" type="button" onclick="SelectAvatar(4)" id="picture_4"><img src="/static/images/4.webp" alt="Profile picture with ID 4" width="200px" height="200px"></button>
      </div>
      <p id="status"></p>
      <button class="w-fit bg-[white] text-[black] text-[JetBrains_Mono] rounded-[5px] m-[10px] p-[5px]" type="button" onclick="SendChangeAvatarRequest(0)" id="change_avatar_button">Change avatar</button>
    `

  const changeavatarmenu = document.createElement("div")
  changeavatarmenu.setAttribute("class", "h-fit w-fit bg-[#333c46] absolute rounded-[5px] p-[10px]")
}

async function SendChangeAvatarRequest(avatar_id) {
  if (avatar_id == 0) {
    document.getElementById("status").textContent = `You don't select avatar!`
    return
  }
  const change_avatar_request = await fetch("/api/profile/avatar", {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      Profile_pic: avatar_id
    })
  })
  if (change_avatar_request.ok) {
    const message = await change_avatar_request.json()
    document.getElementById("status").textContent = message.Message
  }
  else{
    const message = await change_avatar_request.json()
    document.getElementById("status").textContent = message.Message
  }
}
