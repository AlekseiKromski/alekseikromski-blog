import "./post.css"
import ChatBubbleOutlineIcon from '@mui/icons-material/ChatBubbleOutline';
import {Link} from "react-router-dom";
import TiptapView from "../tiptap/tiptapView";

function Post (props) {
    return (
        <Link to={`/post/${props.id}`} className="post">
            <div className="image" style={{
                "backgroundImage": `linear-gradient(to bottom, rgba(252, 252, 252, 0) 20%, rgb(71 71 71) 100%), url('${props.img}')`
            }}></div>
            <div className="postContent">
                <h1>{props.title}</h1>
                <TiptapView content={`<p>${props.description.slice(0, 90)}...</p>`}/>
                <div className="info">
                    <span>{props.createdAt.toString()}</span>
                    <div className="comments">
                        {props.comments}
                        <ChatBubbleOutlineIcon className="commentsIcon"/>
                    </div>
                </div>
            </div>
        </Link>
    )
}

export default Post