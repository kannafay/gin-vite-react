import { NavLink } from 'react-router'

export default function Home() {
  return (<>
    <h1>Home</h1>
    <NavLink to="/about">跳转到关于</NavLink>
  </>)
}