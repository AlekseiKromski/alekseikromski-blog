import {Link} from "react-router-dom";
import "./main.css"
import Post from "../../components/post/post"
import PostMock from "../../components/post-mock/postMock"
import {useEffect, useState} from "react";
import axios from "axios";

function Main() {

    let [loading, setLoading] = useState(true)
    let [posts, setPosts] = useState([])

    useEffect(() => {
        axios.get("http://localhost:3001/v1/post/get-last-posts/15/0").catch(
            setPosts([])
        ).then(response => {
            setPosts(response.data)
        })

        setTimeout(() => {
            setLoading(false)
        }, 1000)

    }, []);

    return (
        <div className="main">
            <h1>Posts</h1>
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
                        <p>No content</p>
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