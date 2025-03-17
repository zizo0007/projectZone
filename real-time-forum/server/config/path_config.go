package config

// Define a base path for templates
var BasePath = "../"

var Container = `
<div class="container">
    <div class="posts">
        <div class="posts-header">
            <button class="nav-button" onclick="displayMobileNav()">
                <i class="fa-solid fa-bars"></i>
            </button>
            <a href="/post/create" class="create-post-link">
                <i class="fa-solid fa-plus"></i>
                Create post
            </a>
        </div>
        {{if .Data}}
        {{range .Data}}
        <div class="post">
            <div class="post-body">
                <a href="/post/{{.ID}}" class="post-title">{{.Title}}</a>
                <div class="post-header">
                    <p class="post-user">{{.UserName}} </p>
                    <span></span>
                    <p class="post-time" data-timestamp="{{.CreatedAt}}">{{.CreatedAt}}</p>
                </div>
                <p class="post-content" id="post-content-home">{{.Content}} </p>
                <div class="post-categories">
                    {{range .Categories}}
                    <span class="post-category">#{{.}}</span>
                    {{end}}
                </div>
            </div>
            <div class="post-footer">
                <button id="likescount{{.ID}}" onclick="postreaction('{{.ID}}','like')"
                    class="post-like post-footer-hover"><i class="fa-regular fa-thumbs-up"></i>{{.Likes}}</button>
                <button id="dislikescount{{.ID}}" onclick="postreaction('{{.ID}}','dislike')"
                    class="post-dislike post-footer-hover"><i
                        class="fa-regular fa-thumbs-down"></i>{{.Dislikes}}</button>
                <a href="/post/{{.ID}}" class="post-comments post-footer-hover">
                    <i class="fa-regular fa-comment"></i>{{.Comments}}
                </a>
            </div>
            <span style="color:red" id="errorlogin{{.ID}}"></span>
        </div>
        {{end}}
        {{else}}
        <p class="no-posts">No posts available to display !</p>
        {{end}}
    </div>
    <div class="pagination">
        <a onclick="pagination('back', '{{if .Data}}true{{end}}')" class="back" href="#">&laquo;
            Back</a>
        <span class="currentpage">1</span>
        <a onclick="pagination('next', '{{if .Data}}true{{end}}')" class="next" href="#">Next
            &raquo;</a>
    </div>
    <script>
        const queryString = window.location.search;
        const urlParams = new URLSearchParams(queryString);
        const path = window.location.pathname
        let page = 1

        if (!isNaN(parseInt(urlParams.get('PageID')))) {
            page = parseInt(urlParams.get('PageID'))
        }
        fetch(path + "?PageID=" + (page + 1)).then(response => {
            if (response.status != 200) {
                const nextbtn = document.querySelector(".next")
                nextbtn.outerHTML = '<a class="next" style="cursor : not-allowed; color : grey;">Next &raquo;</a>'
            }
        })

        if (urlParams.get('PageID') <= 1) {
            const backbtn = document.querySelector(".back")
            backbtn.outerHTML = '<a class="back" style="cursor : not-allowed; color : grey;">&laquo; Back</a>'
        }
        document.querySelector(".currentpage").innerText = urlParams.get('PageID') > 0 ? urlParams.get('PageID') : 1
    </script>
</div>
</div>
`
