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
        <a href="/article/${article.Id}">${article.Title}</a>
        <hr>
        `

      articles_element.appendChild(article_element);
    });
  }
}

export default Fetch_Articles()
