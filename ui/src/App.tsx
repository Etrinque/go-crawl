import './App.css'
import ConfigContainer from "./components/ConfigContainer.tsx";
import Summary from "./components/Summary.tsx";
import LogOutput from "./components/LogOutput.tsx";
import {BiLogoGoLang, BiLogoReact, BiLogoTailwindCss} from "react-icons/bi";

function App() {


  return (
      <div className={""}>
        <span className={"header "}>GO-CRAWL</span>
        <span className={"flex justify-center gap-3 m-2 text-4xl text-center"}>
          <button>
            <a>
              <BiLogoGoLang className={"text-cyan-400/65"}/>
            </a>
          </button>
          <button className={" cursor-pointer"}>
            <a href={"https://reactjs.org/"}>
              <BiLogoReact className={"text-cyan-400/65"}/>
            </a>
          </button>
          <button>
            <a>
              <BiLogoTailwindCss className={"text-cyan-400/65"}/>
            </a>
          </button>
        </span>
        <div className={"app "}>
          {/*  LOGO HERE  */}
          <ConfigContainer/>
          <Summary/>
          <LogOutput/>
        </div>
      </div>
  )
}

export default App
