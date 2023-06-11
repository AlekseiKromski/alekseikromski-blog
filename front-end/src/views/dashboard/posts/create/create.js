import {Link, useNavigate} from "react-router-dom";
import TiptapEdit from "../../../../components/tiptap/tiptapEdit/tiptapEdit";
import {useSelector} from "react-redux";
import {useRef, useState} from "react";
import {useEditor} from "@tiptap/react";
import StarterKit from "@tiptap/starter-kit";
import TextAlign from "@tiptap/extension-text-align";
import Highlight from "@tiptap/extension-highlight";
import Dropcursor from '@tiptap/extension-dropcursor'
import Image from '@tiptap/extension-image'
import styles from "./create.module.css"
import Loading from "../../../../components/loading/loading";

function PostCreate() {
    let upload = useRef(null)
    let navigate = useNavigate()

    let [title, setTitle] = useState("")
    let [image, setImage] = useState("")
    let [category, setCategory] = useState(0)
    let [canSend, setCanSend] = useState(true)
    let [loading, setLoading] = useState(false)

    const application = useSelector((state) => state.application);
    const shared = useSelector((state) => state.shared);

    const editor = useEditor({
        extensions: [
            StarterKit,
            TextAlign.configure({
                types: ['heading', 'paragraph'],
            }),
            Highlight,
            Image,
            Dropcursor,
        ],
        content: "",
    })

    function createPost(){
        if (canSend) {
            application.axios.post("/v1/post/create-post", {
                "title": title,
                "description": editor.getHTML(),
                "img": image,
                "category_id": category
            }).then(() => {
                navigate("/dashboard/admin/posts")
            })
        }
    }

    const addImage = (ref) => {
        let file = ref.current.files[0]
        let form = new FormData()

        form.append("file", file)

        setCanSend(false)
        setLoading(true)
        application.axios.post("/v1/post/upload", form)
            .then(response => {
                setImage(`${process.env.REACT_APP_AXIOS_BASE_URL}${response.data.payload}`)
                setCanSend(true)
                setLoading(false)
            })
            .catch(e => {
                console.error(e)
            })
    }


    return (
        <div className={styles.createMain}>
            <h1><Link to={"/dashboard/admin"}>Dashboard</Link> / <Link to={'/dashboard/admin/posts/'}>Posts</Link> / Create</h1>

            <div className={styles.basicData}>
                <div className="">
                    <label htmlFor="">Title</label>
                    <input type="text" value={title} onChange={(e) => setTitle(e.target.value)}/>
                </div>

                <div className="">
                    <label htmlFor="">Image</label>
                    <input type="file" ref={upload} onChange={() => addImage(upload)}/>
                </div>

                <div className="">
                    <label htmlFor="">Category</label>
                    <select onChange={(e) => setCategory(Number.parseInt(e.target.value))}>
                        {shared.categories != null &&
                            shared.categories.map(category => (
                                <option value={category.ID} key={category.ID} >{category.name}</option>
                            ))
                        }
                    </select>
                </div>
            </div>

            <TiptapEdit setLoading={setLoading} setCanSend={setCanSend} editor={editor}/>

            <button onClick={() => {createPost()}}>
                { loading ?
                    <Loading/>
                    :
                    <span>Create</span>
                }
            </button>
        </div>
    )
}

export default PostCreate