# CAD Mechanical Design Learning Pathway (Associate)
https://www.autodesk.com/certification/learning-pathways/cad-mechanical-design
- https://www.autodesk.com/certification/learn/course/fusion360-intro-to-3d-modeling-associate <- You are here
- https://www.autodesk.com/certification/learn/course/fusion360-intro-design-manufacturing-associate

---

# 2: Introduction to modeling
## Create a new project
- Oh dang didn't even know you had projects and folders
- Orbit and look at are very handy to visualize parts!
- Fusion has a "z is up" default
- You can select objects in the drawing and they will be highlighted in the folder explorer as well as the timeline 
- Right click > find in browser or find in window, neat!

## Create and edit a sketch
- When create an object (ex: circle) options are displayed in the sketch pallete in the right side!
- Oooo vertical constrain on two dots, that's new!
- Hold click left to select something when you have too much going on
- Hold down ctrl to add additional selections
- Oooh the line tool has a checkmark and I never saw it ahahah HELP
- The trim tool can trim away the excess of a line, so cool! Oh wow, you can even trim the circle
- When you wanna snap a line into something and you see the triangle icon, it means you reached the midpoint of the thing. So cool!
- Wow you can reference one angle when creating another angle for a dimmension constrain by selecting it when applying the dimmension constrain

## Create and edit a 3D model 
- Revolve, extrude and fillet, neat!
- Pattern can select features, that's really cool (so you select extrude+chamfer features)
- This disk brake turned out awesome

## Create a technical drawing
- Remember to add multiple views!
- You can have multible base views, noice
- Woah this is really comprehensive. I don't think I'll need technical drawings for my 3d prints, but revisit this chapter in case you do

## Practice exercise
- Used a constraint to center some bolts in a hole, simple but neat

## Challenge exercise
- Used extrude to cut away a bolt head, then a sketch and another extrude to recreate it!

# Introduction to parametric sketching

## Create parameter based sketches
- User parameters, construction planes, circles and rectangles!
- Link sketch dimensions and parameters!

- Offset construction plane, woah
- Looks like you can construct a lot in 2d before doing 3d, I just didn't think about it this way before

## Sketch splines and slots
- We'll create a sketch slot, a spline and a conic

- You can select two points and use the horizontal constrain to make sure they're horizontal with each other, super cool
- 
- Fit point splines are super powerful! They to take a bit of "finesse" to use
- Author recomends to create splines with the least amount of spline control planes as possible
- You can also add additional points to the control point spline
- Every sketch type has additional options when selected on the sketch palette (for splines, you can display the curvature combs)
- Conic curves allow you to create a curve that is more squared, or less squared, depending on the Rho value
- Dang the elipse is really neat

Author encourages you to practice with these elements! They looked cool.

## Sketch text
- Text on a line. Text on a box
- You can explode the text to turn it into shapes on a sketch and modify the letters! Really cool

## Create sketch intersections and projections
- Sketch projections, intersections, and including 3d geometry in a sketch
- You can project a body, neat!
- Intersect creates... an intersect curve on the sketch plane?
- Good heavens you can sketch in 3d

# 4 - Introduction to parametric modeling

## Create a mechanical link
- This video is just so cool
- Uh-oh, how to copy a design with no history without references between the old and the new one
- Did a bunch of sketching and extruding, this instructor knows a lot, I should try re-creating this link sometime by memory

## Add a sketch canvas image
- You can insert an image, and then right click on the canvas and select "calibrate" to size the image to the correct size!
- Oh you can right click while creating a dimension and create an "alligned" dimmension!
- Of course you can apply constrains to the spline orientation tool
- To constrain a spline you can use "fix/unfix" to dimension the spline points!

## Create a 3D model solid trigger
- Extrude has a taper angle, handy for injection molding, how cool
- Fillets. Lots of fillets
- Draft analysis was used... but I don't understand very well which purpose it serves

## Manage physical materials and appearances
- You can drag and drop a material in the material window over the "in this design" picture thingy. Or drag it over the body, or the componenr on the browser
- Right clck material > edit > advanced
- Appearances can override physical materials, even though the physical properties will remain the same
- Right click components > properties 
- Ooo in the appearance menu you can select a face and apply an appearance only to a face!
- The render workspace (chosen in the top left menu box) allows you to see your part with prettier detail! It shows reflections!
- To remove an appearance, under the appearance menu in the "in this design box", right click and "unassign and delete"

## Practice exercise
- Don't forget that sketch fillets are a thing!
- Used the shell command to create a hollow object, I forgot how handy it was
- Select a line and press X: Turns them into construction
- Author dimensions a point relative to the origin, and then dimensions all points to the point that was locked to the origin first, interesting
- You can select multiple features of a slot and the click the equal constraint, neat
- You can create a construction line and constrain points from several figures on it!
- Making lines parallel can be very handy to constrain things!
- This was a great video with many insights on how fusion 360 works
- You can also roll back and apply a feature and later roll forward and keep that feature (when the shell conflicted with the extrude)

# 5 - Introduction to freeform and direct modeling
- Dang this chapter goes really fast and is a bit beyond my skill level. I'm currently more interested in parametric design, I'll save this one for later

# 6 -  Introduction to assembly modeling
## Introduction to assembly modeling
- Create a component, join, edit a joint limit, drive a joint
- Each component has an origin location!
- You can create a component from a body 
- Components can be moved without using the move command, interesting!

- Grouding a component: Coordinates of the componen stay fixed relative to the design (not too sure what this means=
- When you create a rigid group between 2 or more components, when you drag them around they move together!
- You can select 2 components > right click > isolate (this way only these 2 are visible)

- There are 2 main types of joints
    - As build joint: Makes use of the current location of components
    - Joint: Requires you to select all applicable selections to make the joint happen

- Modify: Allign > select 2 faces > faces are now alligned!
    - You can then select the top faces and the trigger structure now matches!
    - This could've been done with a joint but joints can be tricky sometimes!
    - Move/copy would also have accomplished this but it doesn't give you the ability to allign faces

- As build joint note: Whatever you select first out of the two objects, will be the object that movies

- HEY THERE'S A JOINTS FOLDER!

- Wew man I'm getting my ass kicked by these joints, I need to practice a bit on my own I think
    - Practice a joint in place in which you can chose which object slides!

### TODO here:
- Practice joins (which?)
    - See if you can figure out all the as-built-joint types!
- Doublecheck what rigid groups do
- Practice the assemble functionalities?

Course URL: https://www.autodesk.com/certification/learn/course/fusion360-intro-to-3d-modeling-associate