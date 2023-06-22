import {Link, useNavigate, useParams} from "react-router-dom";
import {useEffect, useState} from "react"
import "./single.css"
import SinglePostMock from "../../components/singlePostMock/singlePostMock";
import ArrowBackIosIcon from '@mui/icons-material/ArrowBackIos';
import {useSelector} from "react-redux";
import TiptapView from "../../components/tiptap/tiptapView/tiptapView";
import {useTranslation} from "react-i18next";
import ReCaptcha from "@matt-block/react-recaptcha-v2";
import Alert from "../../components/alert/alert";

function Single() {
    const {t} = useTranslation()
    const navigate = useNavigate()
    const params = useParams()

    //redux
    const application = useSelector((state) => state.application);

    let [loading, setLoading] = useState(true)
    let [post, setPost] = useState(null)
    let [commentName, setCommentName] = useState("")
    let [commentText, setCommentText] = useState("")
    let [captchaToken, setCaptchaToken] = useState(null)
    let [error, setError] = useState(null)

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
        if (commentName != "" && commentText != "" && captchaToken != null) {
            application.axios.post("/v1/post/comment", {
                name: commentName,
                text: commentText,
                post_id: Number.parseInt(params.id),
                captchaToken: captchaToken
            }).then(response => {
                setPost({
                    ...post,
                    comments: [response.data, ...post.comments]
                })

                setCommentName("")
                setCommentText("")
            }).catch(e => {
                setError("Cannot add comment, try again 🤯")
            })
        }
    }

    return (
        <div className={`singlePost ${!application.sideClosed ? "static" : setTimeout(() => "", 1000)}`}>
            <a onClick={() => navigate(-1)}>
                <ArrowBackIosIcon/>
            </a>
            {loading &&
                <SinglePostMock/>
            }
            {!loading && post !== null &&
                <div className="singlePostContent">
                    {error != null &&
                        <Alert
                            title="Error"
                            type="error"
                            text={error}
                            set={setError}
                        />
                    }
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
                    <span className="singleCategory">{t("single.category")}: <span><Link to={`/${post.category_id}`}>{post.category.name}</Link></span></span>
                    {/*<div>*/}
                    {/*    <b>Tags: </b>*/}
                    {/*    {*/}
                    {/*        post.tags.map( tag => (<span className="tag">{tag.name}</span>))*/}
                    {/*    }*/}
                    {/*</div>*/}
                    <div className="commentForm">
                        <h1>{t("single.comment")}</h1>

                        <div className="">
                            <label htmlFor="name">{t("single.input_name")}</label>
                            <input type="text" name="Name" value={commentName} onChange={e => setCommentName(e.target.value)}/>
                        </div>

                        <div className="">
                            <label htmlFor="comment">{t("single.input_comment")}</label>
                            <textarea name="comment" value={commentText} onChange={e => setCommentText(e.target.value)}></textarea>
                        </div>
                        <ReCaptcha
                            siteKey={process.env.REACT_APP_GOOGLE_RECAPTCHA_TOKEN}
                            theme="light"
                            size="normal"
                            onSuccess={(c) => setCaptchaToken(c)}
                            onError={() => setError("There is some troubles, try again")}
                            onExpire={() => setError("Verification has expired, re-verify.")}
                        />
                        {captchaToken != null &&
                            <button onClick={sendComment}>{t("single.send")}</button>
                        }
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