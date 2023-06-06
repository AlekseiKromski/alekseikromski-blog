import {Link} from "react-router-dom";
import "./main.css"
import Post from "../../components/post/post"
import PostMock from "../../components/post-mock/postMock"
import {useState} from "react";

function Main() {
    let [loading, setLoading] = useState(true)
    setTimeout(() => {
        setLoading(false)
    }, 1000)
    const posts = [
        {
            img: "https://www.bsr.org/images/heroes/bsr-focus-nature-hero.jpg",
            title: "Mock post",
            description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
            createdAt: "2023-05-06 13:30:30",
            comments: 12
        },
        {
            img: "https://www.bsr.org/images/heroes/bsr-focus-nature-hero.jpg",
            title: "Mock post",
            description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
            createdAt: "2023-05-06 13:30:30",
            comments: 12
        },
        {
            img: "https://www.bsr.org/images/heroes/bsr-focus-nature-hero.jpg",
            title: "Mock post",
            description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
            createdAt: "2023-05-06 13:30:30",
            comments: 12
        },
        {
            img: "https://www.bsr.org/images/heroes/bsr-focus-nature-hero.jpg",
            title: "Mock post",
            description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
            createdAt: "2023-05-06 13:30:30",
            comments: 12
        },
        {
            img: "https://www.bsr.org/images/heroes/bsr-focus-nature-hero.jpg",
            title: "Mock post",
            description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
            createdAt: "2023-05-06 13:30:30",
            comments: 12
        },
        {
            img: "https://www.bsr.org/images/heroes/bsr-focus-nature-hero.jpg",
            title: "Mock post",
            description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
            createdAt: "2023-05-06 13:30:30",
            comments: 12
        },
        {
            img: "https://www.bsr.org/images/heroes/bsr-focus-nature-hero.jpg",
            title: "Mock post",
            description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
            createdAt: "2023-05-06 13:30:30",
            comments: 12
        },
        {
            img: "https://www.bsr.org/images/heroes/bsr-focus-nature-hero.jpg",
            title: "Mock post",
            description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
            createdAt: "2023-05-06 13:30:30",
            comments: 12
        },
        {
            img: "https://www.bsr.org/images/heroes/bsr-focus-nature-hero.jpg",
            title: "Mock post",
            description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
            createdAt: "2023-05-06 13:30:30",
            comments: 12
        },
    ]
    return (
        <div className="main">
            <h1>Posts</h1>
            {!loading ?
                <div className="post-map">
                    {
                        posts.map(post => {
                            return (
                                <Post post={post} />
                            )
                        })
                    }
                    {posts.length == 0 &&
                        <p>No content</p>
                    }
                </div>
                :
                <div className="post-map">
                    {
                        [...Array(20).keys()].map(index => {
                            return (
                                <PostMock/>
                            )
                        })
                    }
                </div>

            }
        </div>
    );
}

export default Main;