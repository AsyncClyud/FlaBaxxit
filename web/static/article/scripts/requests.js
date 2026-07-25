"use strict";

function GetPathValue() {
  var URL = (window.location.pathname)
  var LastIndexURL = URL.substring(0, URL.length - 1).lastIndexOf("/")
  const Id = URL.substring(LastIndexURL+1, URL.length+1)
  return Id
}
async function FetchArticle() {
  const Id = GetPathValue()
  const get_article_request = await fetch(`/api/articles/${Id}`, {
    method: "GET",
    headers: { Accept: "application/json" },
  });
  if (get_article_request.ok) {
    const article = JSON.parse(await get_article_request.json());
    const article_element = document.getElementById("article");
    article_element.setAttribute("class", "w-[90vw] h-fit bg-[#212121] rounded-[5px]");
    article_element.innerHTML = `
    <h3 class="text-[4vh] m-[10px]" id="title">${article.Title}</h3>
    <p class="text-[2vh] m-[5px] px-8" id="content">${article.Content}</p>
    <p class="text-[gray] m-[20px]" id="created_at">${article.Created_at}</p>
    <p class="m-[5px]" >Article Author ID:</p>
    <p class="m-[20px]" id="author_id">${article.Author}</p>
    <h3 class="h-[min-content] w-[65vw] bg-[#333c46] text-[2vh] rounded-[10px] m-[3px] ml-auto mr-auto outline-[3px] outline-solid outline-[#151b23]">Comments</h3>
    <div class="h-[min-content] w-[60vw] bg-[#2a323c] rounded-[5px] m-[10px] ml-auto mr-auto p-4" id="comments"></div>
    `;
    document.title = article.Title

    await GetArticleAuthor()
  }
  if (document.getElementById("author_id").textContent == "0") {
    window.location.replace("/not_found")
  }
}

async function GetArticleAuthor() {
  const id = GetPathValue()
  const author_id = document.getElementById("author_id").textContent

  const article_author_request = await fetch("/api/users", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      Author: Number(author_id)
      })
  })
  if (article_author_request.ok) {
    const main_element = document.getElementById("main")
    const actions_element = document.createElement("div")
    actions_element.setAttribute("class", "h-fit w-[90vw] bg-[#212121] p-6")
    actions_element.innerHTML = `
      <button class="w-fit bg-[white] text-[black] text-[JetBrains_Mono] rounded-[3px] m-[5px] p-[5px]" type="button" onclick="SendDeleteArticleRequest()">Delete Article</button>
      <button class="w-fit bg-[white] text-[black] text-[JetBrains_Mono] rounded-[3px] m-[5px] p-[5px]" type="button" onclick="UpdatePageRedirect(${id})">Edit Article</button>
      `
    main_element.appendChild(actions_element)
  }
}

async function FetchArticleComments() {
  const Id = GetPathValue()

    const comments_request = await fetch(`/api/comments/${Id}`, {
    method: "GET",
    headers: { "Accept": "application/json" }
  })
  if (comments_request.ok) {
    const comments = JSON.parse(await comments_request.json())
    const comments_element = document.getElementById("comments")
    comments.forEach((comment) => {
      const comment_element = document.createElement("div");
      comment_element.innerHTML = `
        <h3>Author Id: ${comment.Author}</h3>
        <p class="m-[5px]">${comment.Comment_content}</p>
        <p class="text-[gray] m-[20px]">${comment.Created_at}</p>
        <hr>
        `
        comments_element.appendChild(comment_element)
    });
  }
}

async function SendCreateArticleRequest() {
  const title = document.getElementById("title").value
  const content = document.getElementById("content").value
  const turnstile_token = turnstile.getResponse()

  const create_request = await fetch("/api/articles", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      Title: title,
      Content: content,
      Turnstile_token: turnstile_token
      })
  })
  if (create_request.ok) {
    const message = await create_request.json()
    document.getElementById("status").textContent = message.Message
  }
  else {
    const message = await create_request.json()
    document.getElementById("status").textContent = message.Message
    turnstile.reset()
  }
}

async function SendDeleteArticleRequest() {
  const article_id = GetPathValue()
  const author_id = document.getElementById("author_id").textContent

  const delete_request = await fetch("/api/articles", {
    method: "DELETE",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      Id: Number(article_id),
      Author: Number(author_id)
    })
  })
  if (delete_request.ok) {
   window.location.replace("/")
  }
}

async function SendUpdateArticleRequest() {
  const new_title = document.getElementById("new_title").value
  const new_content = document.getElementById("new_content").value
  const author_id = document.getElementById("author_id").textContent
  const turnstile_token = turnstile.getResponse()

  const update_request = await fetch("/api/articles", {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      Id: Number(GetPathValue()),
      Title: new_title,
      Content: new_content,
      Author: Number(author_id),
      Turnstile_token: turnstile_token
    })
  })
  if (update_request.ok) {
    const message = await update_request.json()
    document.getElementById("status").textContent = message.Message
    await new Promise(r => setTimeout(r, 2000));
    window.location.replace(`/article/${Number(GetPathValue())}`)
  }
  else {
    const message = await update_request.json()
    document.getElementById("status").textContent = message.Message
    turnstile.reset()
  }
}

function UpdatePageRedirect(article_id) {
  window.location.replace(`/article/update/${article_id}`)
}

async function SendCreateCommentRequest() {
  const comment_content = document.getElementById("comment_content").value
  const post_id = GetPathValue()
  const turnstile_token = turnstile.getResponse()

  const create_request = await fetch("/api/comments", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      Comment_content: comment_content,
      Post_id: Number(post_id),
      Turnstile_token: turnstile_token
      })
  })
  if (create_request.ok) {
    const message = await create_request.json()
    document.getElementById("status").textContent = message.Message
    await new Promise(r => setTimeout(r, 2000));
    window.location.replace("/profile")
  }
  else {
    const message = await create_request.json()
    document.getElementById("status").textContent = message.Message
    turnstile.reset()
  }
}
