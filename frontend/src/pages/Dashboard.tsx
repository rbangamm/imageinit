import { useContext } from "react";
import { UserContext } from "../hooks/UserContext";

const Dashboard = () => {
    const {setToken} = useContext(UserContext);
    return (
        <div>
            <h1>Dashboard</h1>
            <button>Add Image</button>
            <button>Delete Image</button>
            <button onClick={() => {setToken("invalid")}}>Logout</button>
        </div>
    );
}

export default Dashboard;