import "./App.css";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import Invitation from "./page/invitation";
import RSVP from "./page/rsvp";
import Home from "./page/home";
import AdminEvent from "./page/admin/event";
import AdminAuth from "./page/admin/auth";
import AdminPeople from "./page/admin/people";
import CreateEvent from "./page/admin/event/create";
import CreatePerson from "./page/admin/people/create";
import Invite from "./page/admin/event/invite";

const App = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/invitation/:id" Component={Invitation} />
        <Route path="/invitation/:id/rsvp" Component={RSVP} />
        <Route path="/admin" Component={AdminAuth} />
        <Route path="/admin/event" Component={AdminEvent} />
        <Route path="/admin/event/:id/invite" Component={Invite} />
        <Route path="/admin/people" Component={AdminPeople} />
        <Route path="/admin/event/create" Component={CreateEvent} />
        <Route path="/admin/people/create" Component={CreatePerson} />
        <Route path="/" Component={Home} />
        <Route path="/*" Component={() => <div>Not Found</div>} />
      </Routes>
    </BrowserRouter>
  );
};

export default App;
