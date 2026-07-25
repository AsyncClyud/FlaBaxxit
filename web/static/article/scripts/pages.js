"use strict";

function GetPathValue() {
  var URL = (window.location.pathname)
  var LastIndexURL = URL.substring(0, URL.length - 1).lastIndexOf("/")
  const Id = URL.substring(LastIndexURL+1, URL.length+1)
  return Id
}

async function FetchArticleInfo() {
  const id = GetPathValue()
  const getarticle_info = await fetch(`/api/articles/${id}`, {
    method: "GET",
    headers: { Accept: "application/json" },
  });
  if (getarticle_info.ok) {
    const article_info = JSON.parse(await getarticle_info.json())
    document.getElementById("title").textContent = `${article_info.Title}`
    document.getElementById("new_content").textContent = `${article_info.Content}`
    document.getElementById("author_id").textContent = `${article_info.Author}`
    if (document.getElementById("author_id").textContent == "0") {
      window.location.replace("/not_found")
    }
  }
}

export default GetPathValue(); FetchArticleInfo()
