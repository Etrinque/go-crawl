// type CrawlProps = {
//   url: string
//   maxPages: number
//   numWorkers: number
// }
import {TbInfoSquareRounded} from "react-icons/tb";
import {useState} from "react";

const ConfigContainer = () => {
  // api calls here
  // const crawl: CrawlProps = {url: "https://www.google.com", maxPages: 10, numWorkers: 10}

  const [numWorkers, setWorkers] = useState(0)
  const [maxPages, setMaxPages] = useState(0)
  const [depthTooltipVisible, setDepthTooltipVisible] = useState(false)
  const [workerTooltipVisible, setWorkerTooltipVisible] = useState(false)

  return (
      <div
          className={"card"}>
        <form>
          <fieldset
              className={"flex flex-col  p-2 text-start  [&>input]:bg-neutral-800 [&>input]:border-cyan-400/50  "}>

            <label htmlFor="url" className={" pb-2 font-body text-2xl text-center"}>Root URL </label>
            <input
                className={" border  focus:outline-none focus:border-cyan-400/75 rounded-xl p-2 m-2"}
                placeholder={"https://example.com"}
                type="text"/>

            <div className={"accent-cyan-500/60 p-2 text-xl"}>
              <div className={"flex flex-col gap-3 p-2 items-center"}>
                <span className={"flex gap-2"}>
                  <label htmlFor="maxPages" className={"font-body"}>
                    Crawl Depth {maxPages}
                  </label>

                   <TbInfoSquareRounded className={"relative h-4 w-4 text-cyan-400/60"} onMouseEnter={() => {
                     setDepthTooltipVisible(true)
                   }} onMouseLeave={() => {
                     setDepthTooltipVisible(false)
                   }}/>
                  <div
                      id={"depth-tooltip"}
                      style={{display: depthTooltipVisible ? "block" : "none"}}
                      className={"absolute  text-sm z-auto bg-neutral-300/75 text-neutral-800 p-2 rounded-xl"}>
                    <p className={""}>
                      Maximum pages to hit: 1..1000
                    </p>
                  </div>
                </span>

                <input
                    className={""}
                    id="maxPages"
                    name="maxPages"
                    min={1}
                    max={1000}
                    step={1}
                    value={maxPages}
                    type="range"
                    onChange={(e) => setMaxPages(Number(e.target.value))}
                />
              </div>

              <div className={"flex flex-col gap-2 pt-5 items-center"}>
                <span className={"flex gap-2"}>

                <label htmlFor="numWorkers" className={"font-body"}>
                  Worker Pool {numWorkers}
                </label>
                    <TbInfoSquareRounded
                        className={"relative h-4 w-4 text-cyan-400/60"}
                        onMouseEnter={() => {
                          setWorkerTooltipVisible(true)
                        }}
                        onMouseLeave={() => {
                          setWorkerTooltipVisible(false)
                        }}/>

                <div
                    id={"worker-tooltip"}
                    style={{display: workerTooltipVisible ? "inline-block" : "none"}}
                    className={"absolute cursor-default bg-neutral-300/75 text-neutral-800 p-2 rounded-xl"}>
                  <p className={"text-sm"}>
                    Maximum worker pool size: 1..50
                  </p>
                </div>
                  </span>

                <input
                    value={numWorkers}
                    className={""}
                    id="numWorkers"
                    name="numWorkers"
                    min={1} max={100} step={1}
                    type="range" minLength={100}
                    onChange={(e) => setWorkers(Number(e.target.value))}
                />
              </div>

            </div>


            <div className={"flex justify-center items-center pt-4 font-body"}>
              <label htmlFor="robots.txt" className={"m-2"}>Respect robots.txt?</label>
              <input type={"checkbox"}
                     className={"h-5 w-5 accent-cyan-600 translate-y-[-0.2rem]"}
                     id="robots"/>
            </div>
          </fieldset>

          <div>
            <button className={"form-button "} type={"button"}>Reset</button>
            <button className={"form-button form-button-submit"} type={"button"} typeof={"submit"}>Crawl!</button>
          </div>
        </form>

      </div>
  )
}

export default ConfigContainer;