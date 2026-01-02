import {useState} from "react";

const LogOutput = () => {

  const prompt = "$>"

  const [stream, setStream] = useState<string>(prompt);

  return (
      <div className={"card  flex flex-col gap-4 min-h-[480px]"}>
        <div>

          <span className={"font-heading text-4xl"}> Logs</span>
        </div>

        <div className={"bg-black rounded-xl min-h-[400px]"}>
          <textarea className={"p-4 text-xl w-full h-full min-h-[400px] text-white"} readOnly={true} value={stream}
                    onChange={(e) => setStream(e.target.value)}/>
        </div>


      </div>)
}

export default LogOutput