"use strict";

function GetPathValue() {
  var URL = (window.location.pathname)
  var LastIndexURL = URL.substring(0, URL.length - 1).lastIndexOf("/")
  const Id = URL.substring(LastIndexURL+1, URL.length+1)
  return Id
}

async function FetchUserProfile() {
  const Id = GetPathValue()
  const response = await fetch(`/api/profile/${Id}`, {
    method: "GET",
    headers: { Accept: "application/json" },
  });
  if (response.ok) {
    const data = JSON.parse(await response.json())
    const main_element = document.getElementById("main")
    const user_info = document.createElement("div")
    user_info.setAttribute("id", "profile")
    user_info.innerHTML = `
      <div class="h-fit w-[70vw] bg-[#212121] rounded-[5px]" id="profile">
      <img class="m-[5px] ml-auto mr-auto rounded-full" src="/static/images/${data.Profile_pic}.webp" alt="Avatar" width="150" height="150">
      <h3 class="text-[4vh] m-[10px]" id="username">${data.Username}</h3>
      <p class="text-[2vh] m-[5px] px-8">${data.Bio}</p>
      <p class="text-[gray] m-[20px]">${data.Created_at}</p>
      </div>
    `
    main_element.appendChild(user_info)
    document.title = `${data.Username} profile`
  }
  if (response.status == 401) {
    window.location.replace("/auth/login")
  }
}

export default FetchUserProfile()
