import "./post.css"
import ChatBubbleOutlineIcon from '@mui/icons-material/ChatBubbleOutline';

function Post (props) {
    return (
        <div className="post">
            <div className="image" style={{
                "backgroundImage": `linear-gradient(to bottom, rgba(252, 252, 252, 0) 20%, rgb(71 71 71) 100%), url('${props.post.img}')`
            }}></div>
            <div className="postContent">
                <h1>{props.post.title}</h1>
                <p>{props.post.description.slice(-90)}...</p>

                <div className="info">
                    <span>{props.post.createdAt.toString()}</span>
                    <div className="comments">
                        {props.post.comments}
                        <ChatBubbleOutlineIcon className="commentsIcon"/>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default Post