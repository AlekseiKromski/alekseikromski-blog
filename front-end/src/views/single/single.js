import {Link, useParams} from "react-router-dom";
import {useEffect, useState} from "react"
import axios from "axios";
import "./single.css"
import SinglePostMock from "../../components/singlePostMock/singlePostMock";
import ArrowBackIosIcon from '@mui/icons-material/ArrowBackIos';

function Single() {
    const params = useParams()

    let [loading, setLoading] = useState(true)
    let [post, setPost] = useState(null)

    useEffect(() => {
        axios.get(`http://localhost:3001/v1/post/get-post/${params.id}`).catch(
            setPost(null)
        ).then(response => {
            console.log(response.data)
            setPost(response.data)
        })

        setTimeout(() => {
            setLoading(false)
        }, 1000)

    }, [params.id]);

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
                    <p>{post.description}</p>
                    <span className="singleCategory">Category: <span>{post.category.name}</span></span>
                    <div>
                        <b>Tags: </b>
                        {
                            post.tags.map( tag => (<span className="tag">{tag.name}</span>))
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