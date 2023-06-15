import styles from "./edit.module.css"
import PostCreate from "../create/create";
import {useParams} from "react-router-dom";
import {useEffect, useState} from "react";
import {useSelector} from "react-redux";

function PostEdit() {
    const params = useParams()
    const application = useSelector((state) => state.application);

    let [post, setPost] = useState(null)

    useEffect(() => {
        if (params.id) {
            application.axios.get(`/v1/post/get-post/${params.id}`).catch(
                setPost(null)
            ).then(response => {
                setPost(response.data)
            })
        }
    },[])

    return (
        <div>
            {post == null ?
                <p>No data</p>
                :
                <PostCreate post={post}/>
            }
        </div>
    )
}

export default PostEdit