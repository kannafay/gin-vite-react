import { NavLink, useLoaderData } from 'react-router'

export default function About() {
  const { title } = useLoaderData()

  return (<>
    <h1>{title}</h1>
    <NavLink to="/">跳转到首页</NavLink>
  </>)
}