import { EditorContent } from '@tiptap/react'
import {useRef} from "react";
import {useSelector} from "react-redux";
import styles from "./tiptapEdit.module.css"
import "./tiptap.css"
const TiptapEdit = ({editor, setCanSend, setLoading}) => {
    let upload = useRef(null)

    const application = useSelector((state) => state.application);

    if (!editor) {
        return null
    }

    const addImage = () => {
        let file = upload.current.files[0]
        let form = new FormData()

        form.append("file", file)
        setCanSend(false)
        setLoading(true)
        application.axios.post("/v1/post/upload", form)
            .then(response => {
                editor.chain().focus().setImage({ src: `${process.env.REACT_APP_AXIOS_BASE_URL}${response.data.payload}` }).run()
                setCanSend(true)
                setLoading(false)
            })
            .catch(e => {
                console.error(e)
            })
    }

    return (
        <div className={styles.tiptap}>
            <input className={styles.invisible} onChange={() => addImage()} ref={upload} type="file"/>
            <div className={styles.buttons}>
                <button onClick={() => editor.chain().focus().toggleHeading({ level: 1 }).run()} className={editor.isActive('heading', { level: 1 }) ? 'is-active' : ''}>
                    h1
                </button>
                <button onClick={() => editor.chain().focus().toggleHeading({ level: 2 }).run()} className={editor.isActive('heading', { level: 2 }) ? 'is-active' : ''}>
                    h2
                </button>
                <button onClick={() => editor.chain().focus().toggleHeading({ level: 3 }).run()} className={editor.isActive('heading', { level: 3 }) ? 'is-active' : ''}>
                    h3
                </button>
                <button onClick={() => editor.chain().focus().setParagraph().run()} className={editor.isActive('paragraph') ? 'is-active' : ''}>
                    paragraph
                </button>
                <button onClick={() => editor.chain().focus().toggleBold().run()} className={editor.isActive('bold') ? 'is-active' : ''}>
                    bold
                </button>
                <button onClick={() => editor.chain().focus().toggleItalic().run()} className={editor.isActive('italic') ? 'is-active' : ''}>
                    italic
                </button>
                <button onClick={() => editor.chain().focus().toggleStrike().run()} className={editor.isActive('strike') ? 'is-active' : ''}>
                    strike
                </button>
                <button onClick={() => editor.chain().focus().toggleHighlight().run()} className={editor.isActive('highlight') ? 'is-active' : ''}>
                    highlight
                </button>
                <button onClick={() => editor.chain().focus().setTextAlign('left').run()} className={editor.isActive({ textAlign: 'left' }) ? 'is-active' : ''}>
                    left
                </button>
                <button onClick={() => editor.chain().focus().setTextAlign('center').run()} className={editor.isActive({ textAlign: 'center' }) ? 'is-active' : ''}>
                    center
                </button>
                <button onClick={() => editor.chain().focus().setTextAlign('right').run()} className={editor.isActive({ textAlign: 'right' }) ? 'is-active' : ''}>
                    right
                </button>
                <button onClick={() => editor.chain().focus().setTextAlign('justify').run()} className={editor.isActive({ textAlign: 'justify' }) ? 'is-active' : ''}>
                    justify
                </button>
                <button onClick={() => upload.current.click()}>add image</button>
            </div>
            <EditorContent editor={editor} />
        </div>
    )
}

export default TiptapEdit