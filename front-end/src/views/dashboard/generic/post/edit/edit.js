import styles from "./edit.module.css"
import PostCreate from "../create/create";
import {useParams} from "react-router-dom";
import {useEffect, useState} from "react";
import {useSelector} from "react-redux";
import {useTranslation} from "react-i18next";

function PostEdit() {
    let {t} = useTranslation()
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
                <p>{t("post_update.no_data")}</p>
                :
                <PostCreate post={post}/>
            }
        </div>
    )
}

export default PostEdit