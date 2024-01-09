import "./App.css";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import Invitation from "./page/invitation";
import RSVP from "./page/rsvp";

const App = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/invitation/:id" Component={Invitation} />
        <Route path="/invitation/:id/rsvp" Component={RSVP} />
        <Route path="/*" Component={() => <div>Not Found</div>} />
      </Routes>
    </BrowserRouter>
  );
};

export default App;
