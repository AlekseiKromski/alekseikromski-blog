import {Link, useNavigate} from "react-router-dom";
import TiptapEdit from "../../../../../components/tiptap/tiptapEdit/tiptapEdit";
import {useSelector} from "react-redux";
import {useEffect, useRef, useState} from "react";
import {useEditor} from "@tiptap/react";
import StarterKit from "@tiptap/starter-kit";
import TextAlign from "@tiptap/extension-text-align";
import Highlight from "@tiptap/extension-highlight";
import Dropcursor from '@tiptap/extension-dropcursor'
import Image from '@tiptap/extension-image'
import styles from "./create.module.css"
import Loading from "../../../../../components/loading/loading";
import BreadCrumbs from "../../../../../components/bread-crumbs/bread-crumbs";
import Alert from "../../../../../components/alert/alert";
import {useTranslation} from "react-i18next";

function PostCreate({post}) {
    let {t} = useTranslation()
    let upload = useRef(null)
    let navigate = useNavigate()

    let [title, setTitle] = useState("")
    let [image, setImage] = useState("")
    let [category, setCategory] = useState(0)
    let [canSend, setCanSend] = useState(true)
    let [loading, setLoading] = useState(false)
    let [error, setError] = useState(null)

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
        content: function () {
            // import description if exists
            if (post !== undefined && post !== null) {
                return post.description
            }
            return ""
        }(),
    })

    function updatePost(){
        if (canSend) {
            application.axios.post("/v1/post/edit", {
                id: post.id,
                "title": title,
                "description": editor.getHTML(),
                "img": image,
                "category_id": category
            }).then(() => {
                navigate("/dashboard/admin/posts")
            }).catch(e => {
                setError(e.response.data.message)
            })
        }
    }

    function createPost(){
        if (canSend) {
            application.axios.post("/v1/post/create-post", {
                "title": title,
                "description": editor.getHTML(),
                "img": image,
                "category_id": category
            }).then(() => {
                navigate("/dashboard/admin/posts")
            }).catch(e => {
                setError(e.response.data.message)
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
                setError(e.response.data.message)
            })
    }

    useEffect(() => {
        //init if we have post
        if (post != undefined || post != null) {
            setTitle(post.title)
            setCategory(post.category_id)
            return
        }

        if (shared.categories.length === 1) {
            setCategory(shared.categories[0].id)
        }
    }, [shared.categories])


    return (
        <div className={styles.createMain}>
            {error != null &&
                <Alert
                    title="Error"
                    type="error"
                    text={error}
                    set={setError}
                />
            }

            <BreadCrumbs
                breadcrubms={{
                    title: function (){
                        if (post != null || post !== undefined) {
                            return t("post_create.breadcumbs.update")
                        }
                        return t("post_create.breadcumbs.create")
                    }(),
                    links: [
                        {
                            title: t("generic.posts.breadcumbs_link_title"),
                            link: "/dashboard/admin"
                        },
                        {
                            title: t("generic.posts.breadcumbs_title"),
                            link: "/dashboard/admin/posts"
                        },
                    ],
                    fastActions: []
                }}
            />

            <div className={styles.basicData}>
                <div className="">
                    <label htmlFor="">{t("post_create.form.title")}</label>
                    <input type="text" value={title} onChange={(e) => setTitle(e.target.value)}/>
                </div>

                <div className="">
                    <label htmlFor="">{t("post_create.form.image")}</label>
                    <input type="file" ref={upload} onChange={() => addImage(upload)}/>
                </div>

                <div className="">
                    <label htmlFor="">{t("post_create.form.category")}</label>
                    <select value={category} onChange={(e) => {
                        setCategory(Number.parseInt(e.target.value))
                    }}>
                        {shared.categories != null &&
                            shared.categories.map(category => {
                                return (<option value={category.id} key={category.id} >{category.name}</option>)
                            })
                        }
                    </select>
                </div>
            </div>

            <TiptapEdit setLoading={setLoading} setCanSend={setCanSend} editor={editor}/>

            <button onClick={() => {
                if (post != null && post != undefined){
                    updatePost()
                }else {
                    createPost()
                }
            }}>
                { loading ?
                    <Loading/>
                    :
                    post != null && post != undefined ?
                        <span>{t("post_create.breadcumbs.update")}</span>
                        :
                        <span>{t("post_create.breadcumbs.create")}</span>
                }
            </button>
        </div>
    )
}

export default PostCreate