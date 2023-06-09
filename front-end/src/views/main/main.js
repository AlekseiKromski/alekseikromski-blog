import {useParams} from "react-router-dom";
import "./main.css"
import Post from "../../components/post/post"
import PostMock from "../../components/post-mock/postMock"
import {useEffect, useState} from "react";
import axios from "axios";
import {useSelector} from "react-redux";

function Main() {
    const {categoryID} = useParams()

    //Redux
    const shared = useSelector((state) => state.shared);
    const application = useSelector((state) => state.application);

    let [loading, setLoading] = useState(true)
    let [posts, setPosts] = useState([])

    async function getPostsByCategory(category) {
        setLoading(true)
        await application.axios.get(`/v1/post/get-posts-by-category/${category}/15/0`).catch(
            setPosts([])
        ).then(response => {
            setPosts(response.data)
        })

        setTimeout(() => {
            setLoading(false)
        }, 500)
    }

    async function getPosts(){
        setLoading(true)
        await application.axios.get("/v1/post/get-last-posts/15/0").catch(
            setPosts([])
        ).then(response => {
            setPosts(response.data)
        })

        setTimeout(() => {
            setLoading(false)
        }, 500)
    }

    useEffect(() => {
        let category = "all"
        if (categoryID != undefined) {
            category = Number.parseInt(categoryID)
            if (Number.isNaN(category)) {
                category = "all"
            }
        }

        if (category == "all") {
            getPosts()
        } else {
            getPostsByCategory(category)
        }

    }, [categoryID]);

    return (
        <div className={`main ${!application.sideClosed ? "static" : ""}`}>
            <div className="mainHeader">
                <h1>Posts</h1>
                <select onChange={(e) => {getPostsByCategory(e.target.value)}} name="categoryID" id="">
                    {shared.categories.length != 0 &&
                        shared.categories.map(category => {
                            if (categoryID == category.ID) {
                                return (
                                    <option
                                        selected={true}
                                        value={category.ID}
                                        key={category.ID}
                                    >{category.name}</option>
                                )
                            } else {
                                return (
                                    <option
                                        value={category.ID}
                                        key={category.ID}
                                    >{category.name}</option>
                                )
                            }
                        })
                    }
                </select>
            </div>
            {!loading ?
                <div className="post-map">
                    {
                        posts.map(post => {
                          return (
                              <Post
                                  key={post.id}
                                  id={post.id}
                                  title={post.title}
                                  img={post.img}
                                  description={post.description}
                                  createdAt={post.createdAt}
                                  comments={post.comments.length}
                                  tags={post.tags}
                              ></Post>
                          )
                        })
                    }
                    {posts.length == 0 &&
                        <div className="noContent">
                            <p>🤯 No content</p>
                            <a onClick={() => getPosts()}>Back</a>
                        </div>
                    }
                </div>
                :
                <div className="post-map">
                    {
                        [...Array(20).keys()].map(index => {
                            return (
                                <PostMock key={index}/>
                            )
                        })
                    }
                </div>

            }
        </div>
    );
}

export default Main;