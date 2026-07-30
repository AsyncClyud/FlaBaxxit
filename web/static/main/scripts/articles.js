"use strict";

async function Fetch_Articles() {
  const response = await fetch("/api/articles", {
    method: "GET",
    headers: { Accept: "application/json" },
  });
  if (response.ok) {
    const data = JSON.parse(await response.json());
    const articles_element = document.getElementById("articles");
    data.forEach((article) => {
      const article_element = document.createElement("div")
      article_element.setAttribute("id", `${article.Id}`)
      article_element.innerHTML = `
        <div class="flex flex-col text-left m-[5px]">
          <img class="rounded-full" src="/static/images/${article.Author_Avatar}.webp" alt="Profile Picture" width="45px", height="45px">
          <a class="p-[5px] text-[17px]" href="/profile/${article.Author_Id}">${article.Author_Username}</a>
        </div>
        <a class="text-[20px]" href="/article/${article.Id}"><strong>${article.Title}</strong></a>
        <hr>
        `

      articles_element.appendChild(article_element);
    });
  }
}

export default Fetch_Articles()
