import {Link} from "react-router-dom";
import {useEffect, useState} from "react";
import {useSelector} from "react-redux";
import styles from "./posts.module.css"

function Posts() {

    const application = useSelector((state) => state.application);
    let [posts, setPosts] = useState([])

    async function getPosts(){
        await application.axios.get("/v1/post/get-last-posts/15/0").catch(
            setPosts([])
        ).then(response => {
            setPosts(response.data)
        })
    }

    function deletePost(id) {
        application.axios.get(`/v1/post/delete/${id}`)
            .then(() => {
                setPosts(posts.filter(post => {
                    if (post.id !== id) {
                        return post
                    }
                }))
            })
    }

    useEffect(() => {
        getPosts()
    }, [])

    return (
        <div className={styles.posts}>
            <h1><Link to={"/dashboard/admin"}>Dashboard</Link> / Posts</h1>

            <div className={styles.fastActions}>
                <Link to="/dashboard/admin/posts/create">Create post</Link>
            </div>

            <table>
                <thead>
                    <tr>
                        <th>ID</th>
                        <th>Name</th>
                        <th>Action</th>
                    </tr>
                </thead>
                <tbody>
                {posts != null &&
                    posts.map(post => (
                        <tr>
                            <th>{post.id}</th>
                            <th className={styles.title}>
                                <Link to={"/post/" + post.id}>{post.title}</Link>
                            </th>
                            <th>
                                <div className={styles.action}>
                                    <button>edit</button>
                                    <button onClick={() => {deletePost(post.id)}}>delete</button>
                                </div>
                            </th>
                        </tr>
                    ))
                }
                </tbody>
            </table>
        </div>
    )
}

export default Posts