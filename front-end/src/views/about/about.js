import { Link } from "react-router-dom";

function About() {
    return (
        <div className="about">
            <h1>Hi, my name is Aleksei Kromski</h1>
            <p>This is blog about my IT life. There you can check all my project and comment some posts</p>
            <Link to="/">Read posts</Link>
        </div>
    );
}

export default About;