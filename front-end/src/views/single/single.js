import {Link, useParams} from "react-router-dom";
import {useEffect, useState} from "react"
import axios from "axios";
import "./single.css"
import SinglePostMock from "../../components/singlePostMock/singlePostMock";
import ArrowBackIosIcon from '@mui/icons-material/ArrowBackIos';
import {useSelector} from "react-redux";
import TiptapView from "../../components/tiptap/tiptapView";

function Single() {
    const params = useParams()

    //redux
    const application = useSelector((state) => state.application);

    let [loading, setLoading] = useState(true)
    let [post, setPost] = useState(null)
    let [commentName, setCommentName] = useState("")
    let [commentText, setCommentText] = useState("")

    useEffect(() => {
        application.axios.get(`/v1/post/get-post/${params.id}`).catch(
            setPost(null)
        ).then(response => {
            setPost(response.data)
        })

        setTimeout(() => {
            setLoading(false)
        }, 1000)

    }, [params.id]);

    function sendComment() {
        if (commentName != "" && commentText != "") {
            application.axios.post("/v1/post/comment", {
                name: commentName,
                text: commentText,
                post_id: Number.parseInt(params.id)
            }).then(response => {
                setPost({
                    ...post,
                    comments: [response.data, ...post.comments]
                })

                setCommentName("")
                setCommentText("")
            })
        }
    }

    return (
        <div className="singlePost">
            <Link to="/">
                <ArrowBackIosIcon/>
            </Link>
            {loading &&
                <SinglePostMock/>
            }
            {!loading && post !== null &&
                <div className="singlePostContent">
                    <div className="singleHeader" style={{
                        "backgroundImage": `linear-gradient(to bottom, rgba(252, 252, 252, 0) 20%, rgb(71 71 71) 100%), url('${post.img}')`
                    }}>
                        <h1>
                            {post.title}
                        </h1>
                        <small>{post.createdAt}</small>
                    </div>
                    <p>
                        <TiptapView content={post.description}/>
                    </p>
                    <span className="singleCategory">Category: <span><Link to={`/${post.category_id}`}>{post.category.name}</Link></span></span>
                    <div>
                        <b>Tags: </b>
                        {
                            post.tags.map( tag => (<span className="tag">{tag.name}</span>))
                        }
                    </div>
                    <div className="commentForm">
                        <h1>Comment</h1>

                        <div className="">
                            <label htmlFor="name">Name</label>
                            <input type="text" name="Name" value={commentName} onChange={e => setCommentName(e.target.value)}/>
                        </div>

                        <div className="">
                            <label htmlFor="comment">Comment</label>
                            <textarea name="comment" value={commentText} onChange={e => setCommentText(e.target.value)}></textarea>
                        </div>
                        <button onClick={sendComment}>send</button>
                    </div>
                    <div className="singleComments">
                        {
                            post.comments.map(comment => (
                                <div key={comment.id} className="singleComment">
                                    <div>
                                        <span>{comment.name}</span>
                                        <small>{comment.createdAt}</small>
                                    </div>
                                    <p>{comment.text}</p>
                                </div>
                            ))
                        }
                    </div>
                </div>
            }
        </div>
    )
}

export default Single