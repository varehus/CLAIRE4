(printf("Hello CLAIRE3, this is our init.cl file\n"))

// Mac version
*where* :: "/Users/ycaseau/claire/v4.0/go3"                      // where the init file is
*output* :: "/Users/ycaseau/claire/v4.0/go3/src"
*meta* :: "/Users/ycaseau/Dropbox/src/clairev4.0/src/meta"         // source files on dropbox
*compile* :: "/Users/ycaseau/Dropbox/src/clairev4.0/src/compile"   // source files on dropbox
*bsrc* :: "/Users/ycaseau/claire/v4.0/go/bsrc"
*tsrc* :: "/Users/ycaseau/claire/v4.0/go/test"

// these are the global variables expected by the compiler
RELEASE:float :: 0.02    // August 4th, 2021 

// ***************************************************************************
// *    Part 1: Modules & compiler environment                               *
// ***************************************************************************

(when c := value("compiler") in 
   (c.safety := 5,
    source(c) := "/Users/ycaseau/claire/v4.0/go3/src"))

 // debug test : place methods that  are called with g_test
begin(Core) 

 // very useful
claire/foo(r:relation,x:any) : set
 -> (let r2 := get(inverse, r) in
       case r2
        (table let v := r2[x] in (if (r2.multivalued? != false) (v as set) else set(v)),
         property let v := get(r2, x) in
                    (if (r2.multivalued? != false) (v as set) else set(v)),
         any case r
             (property (if (r.multivalued? != false) { z in r.domain | x %t get(r, z)}
                       else { z in r.domain | get(r, z) = x}),
              table (if (r.multivalued? != false) { z in r.domain | x %t r[z]}
                     else { z in r.domain | r[z] = x}))))

end(Core)


// ***************************************************************************
// *    Part 2: Performance test modules                                     *
// ***************************************************************************

// these are the performance test files of 2020
mFib :: module(part_of = claire,
              source = *tsrc*,
              uses = list(Reader,mClaire),
              made_of = list("testFib"))

mList :: module(part_of = claire,
              source = *tsrc*,
              uses = list(Reader,mClaire),
              made_of = list("testList"))

mSet :: module(part_of = claire,
              source = *tsrc*,
              uses = list(Reader,mClaire),
              made_of = list("testSet"))

mDict :: module(part_of = claire,
              source = *tsrc*,
              uses = list(Reader,mClaire),
              made_of = list("testDict"))

mObj :: module(part_of = claire,
              source = *tsrc*,
              uses = list(Reader,mClaire),
              made_of = list("testObj"))

mSend :: module(part_of = claire,
              source = *tsrc*,
              uses = list(Reader,mClaire),
              made_of = list("testSend"))

mCopy :: module(part_of = claire,
              source = *tsrc*,
              uses = list(Reader,mClaire),
              made_of = list("testCopy"))

// ***************************************************************************
// *    Part 3: Bugs for CLAIRE                                              *
// ***************************************************************************


// parsing bugs: things that cannot get read right
bu1 :: module( uses = list(Reader), source = *bsrc*,
               made_of = list("bstub","bug1"))

// array related bugs
bu2 :: module( uses = list(Reader), source = *bsrc*, 
               made_of = list("bstub", "bug2"))

// table related bugs
bu3 :: module( uses = list(Reader), source = *bsrc*,
               made_of = list("bstub", "bug3"))

// iteration of a union (interpreted) and other patterns
bu4 :: module( uses = list(Reader), source = *bsrc*,
               made_of = list("bstub", "bug4"))

// bugs with floats
bu5 :: module( uses = list(Reader), source = *bsrc*,
               made_of = list("bstub", "bug5"))

// bugs with class & method definitions
bu6 :: module( uses = list(Reader), source = *bsrc*,
               made_of = list("bstub", "bug6"))

// bugs with worlds
bu7 :: module( uses = list(Reader), source = *bsrc*,
               made_of = list("bstub", "bug7"))

// bugs with instantiation & primitive types
bu8 :: module( uses = list(Reader), source = *bsrc*,
               made_of = list("bstub", "bug8"))

// reversible cells from CLP, untyped version
bu9 :: module( uses = list(Reader), source = *bsrc*,
               made_of = list("bstub", "bug9"))

// reversible cells from CLP, typed version
bu10 :: module( uses = list(Reader), source = *bsrc*,
               made_of = list("bstub", "bug10"))

// famous examples (stack example, doc examples ...)
bu11 :: module( uses = list(Reader), source = *bsrc*,
                made_of = list("bstub", "bug11"))

// bug with tuples, lists and sets
bu12 :: module( uses = list(Reader), source = *bsrc*,
                made_of = list("bstub","bug12"))

// test file for handling unknown & inverses
bu13 :: module( uses = list(Reader), source = *bsrc*,
                made_of = list("bstub", "bug13"))

// sudoku example : need to put in the doc - good example of rules & branch
bu14 :: module( uses = list(Reader), source = *bsrc*,
                made_of = list("bstub", "sudoku"))

// compilation bugs (bu14 and bu15) are not tested with claire1

// ***************************************************************************
// *    Part 4: Simple rule examples                                              *
// ***************************************************************************

// these are the old non-regression tests files (refreshed in July 2021)
(printf("Done. \n"))
