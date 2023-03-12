// *******************************************************************************************
// *     log file : started on August 6th, 2021                                              *
// *******************************************************************************************

// THIS LOG FILE sits in the claire/dropbox/compile directory : how we create CLAIRE 2 (our first interpreter)
// from Kernel go files and compiled CLAIRE files (thanks to 3.5 + trans)
// the first step is trans !

shared context:
- developped on Mac / will need to synchronize on cloud to save
- use Visual Studio Editor

// what I need to do

8/7/2021  : start !
- install claire file for Optimize + Generate (copied from trans)
- make sure that claire1 has properly documented modules : claire, mClaire, iClaire
- created init.cl in /go2 with Optimize and Generate
- load Optimize ! => one file at a time :)

(1) osystem.cl
 - get rid of previous status
 - move g_throw and status into osystem.cl

 // A BIG QUESTION: should Optimize produce to_CL and to_C since the go compiler does the job ? 
 // ========> we start with option 1 (Occam's razor) -> get rid of to_C and to_CL
However, we just comment the stuff out ... we could change our mind

8/8/01 : continue with ocall.cl :)
- sugar? in self_print @ Call is too liberal (call.cl)
- Optimize loads !
- Generate load !

8/9/2001 : start Compile(Core)
- great: compile under debugger => see where the problem are !
- NEW : if verbose() := -1 => we do not print the error
Note: this will require a lot of small fixes so that claire1 has the same scope as claire !
(1) method.cl 

8/10/2021 : keep fixing !
-> reached line 100 :)
-> next is add_slot ! (in 3 but not 4)

8/11/2021
-> method.cl compiled ! added a set test to bug12.cl :)
-> need to fix & test trace 
next: see why externC("toto") produces a Cast to obj !!!

8/12/2021
- Aha ! add slot adds a defaut to list-valued  and set-valued slots (avoids crash)
- bug monstrueux avec pretty_print(%.restrictions)
- it took one full day (a) make sure that error are trapped
                      (b) restore global vars pretty.index/pretty/pbreak
  This code, over two files : function.cl and pretty.cl is UGLY - three methods with catch: print@any (func)
  apply_self_print(function.cl)  and printbox/printl 

8/13/2021
-> we are now in types.cl @ the end :)
-> it took for ever to catch a | x  where the white space is wrong => we need to recognize this special char!
Compile(Core) is OK !
- move to Language & Reader
- done !

8/14/2021 : start compiling the generated code
- make_function("hhh") must take an argument (arity)
- adds set_arity, used by c_code(self:Defmethod)
- MakeFunction0 will be used for modules ? 

8/15
- adds two slots / two global variables to the go_producer : cross_safe and cross_throw + produce an error when
   we discover that that the status would change (avoid finding errors in the compiled code)

THE BIG ISSUE : how to fix 
==========================
(1) fix can_thow? when compiling claire1 so that status(m) is right
(2) claire1 compiling should produce a BadMethods list (when EID compiling was used while status is 0)
(3) the claire2 BadList will be smaller, and mostly due to compiler error

WE NOW MOVE TO FIX CLAIRE 1 without any bad methods !
- initialize => insert_definition => close@method
- done : claire 1 is consistent

=> fix-to-compile Core

8/16 + 8/17
 fix Pattern membership  (1 .. X) % ..[tuple(integer,integer)]
 BIG DECISION: end the extensibility for type lattice, keep it for:
   (a) Type_expression = type or Pattern
   (b) collections
   hence > % is extensible
   => <=t and %type is not
   Pattern is a new type_expression

Type expressions are handled in method
   - vmatch extends %t
   - tmatch extends <=t
   - glb is extended with itself
  
// 8/18: Claire1 works :)
- attention : jito => trace does not work
- added VARIANT = false that will be replaced in define.cl by the jito?() test
- fixed ? system  ...
- resume testing compile(Core) 

// 8/19: resume working on compile(Core) => get rid of BAD METHODS (20 !)
  - this implies using NoErrorOptimize ... however, the compiler should apply can_throw on the optimized code !
  - to make open calls safe, we need to be sure that there always exist a restriction that will apply !
    example : glb, less? who are defined on type_expression

KEY POINT: in CLAIRE 4, integer.ident? is false (normal, Equal is needed on a ClaireInteger) however at compile time, 
if g_sort is integer, identification is expected !

next: get_index2 => object.cl
  integer compilation of * fails 
  put it in debug zone ...

 8/22 - resume after break
 - *@(int,int) does not raise an error (safety off => no overflow)
 - AHA ! tuple is not a subclassof list because tuple(X,Y) is a type that is NOT list(X,Y)
     => two consequences : 
       (a) tuple is a bag, not a list in cReflect.go - even if implemented with ClaireList
       (b) there is no inheritance => nth@tuple is needed, nth_bag and length (duplicate list)
 AND COMPILE(CORE) now produce 0 bad methods, move to Language ! 
    
8/23: I need to clean up open !!!!!
// -2   forward       trick used in define.cl to note the forward definition
// -1   closed        c: no more instance, no subclass (but with instances)
// 0   abstract       c: no instances                    p: super closed
// 1   final          c: thing with no subclasses        p: compiled
// 2   default        c: default = ephemeral             p: default
// 3    open          c: keep instances                  p: extensible    
AHA : look for p.open = 4, seems a special marker (compiled)
KEY : interpreted properties have open = 2, when compiled, move to 1.


8/24: close until the book is gone -------------------------------
WHEN I RESUME, Compile(Core) OK, move to Compile(Language)
- compile(Language) OK !
- dn @ integer :

AHA: // iClaire/typing :: Kernel/typing commented out because "this cannot work in CLAIRE 4"   => THIS IS NOT ACCEPTABLE
we need to implement namespaces in CLAIRE 4, this is too convenient

8/28 : resume
HUGE DECISION => keep a safe copy of trans !
For a thing A/x:B (defined in B, in namespace A) we change the go identifier:
     B.C_A_x   if A != claire
     B.C_x     if A = claire

reflect in goexample.cl  :)
rename
  - cident => cap_ident  (capitalized ident, new: adds namespace if not claire)
  - c_ident => go_var  (the go identifier C_...)

=> back to fixing the compiling of claire2 from claire1 :)

- uniformity for methods does not see to include range (look at Copy()) ! > fix goMethod
- abs @ integer (function) is poorly compiled <  -> Reader.F_sup !

8/30: resume !
-> fixed bugs in function.go  (mostly due to inf/sup at integer)
  decisions: (1) use functions and not methods for ClaireFunction

8/31 -> 2/9 : lots of issues
(1) known? @ table should have an inline definition !
    compiler.inline? is set to true => recompile claire1
(2) string should have functions ! .Length() -> F_length_
(3) set! @ list defined in Core ! not in Kernel => function
(4) append @ list => /+ => function !
(5) list! @ set defined in Core => function

9/4 - decision about imported functions
- claire1: avoid the problem, use METHODSET (e.g. list ! @ set)
- claire2: put the # marker on imported functions  (modif in odefine.cl when extract_status returns "body")
- claire3: compile function call while taking the marker into account

-> compile(Core) works ! two more module to go
compile(Language)
  - attach_comment @ any may produce an error, its EID form misses this.

9/6 - g_throw should only be applied to optimized code
  (1) because optimization can remove error-generating ambiguity
  (2) we still need Compile/NoErrorOptimize in osystem.cl (because macro expansion may remove error-generating)
  (3) however write(p,x,y) should NOT be excluded !
  (4) when a method is in this list it does not change its status (avoid generating .ToEid)

=> hence we need to restart with Core ... and see what happens

9/11 - restart after a busy week
- two small bugs and Language is done :)  (go compile OK)
- actually, there are many Kernel method whose status is not properly setup !
- in odefine.cl => disable the code that adds a new object from c to c.instances.
  this is done by new(c) or new(c).is(c)

9/14 - big topic !
iteration must preserve order for lists or intervals ... hence 
- enemerate should return a list, not a set => defined in type.cl, range was bag, must be list
- optimize for over sets in the interpreted code to avoid set to list iteration !
- compilation pattern: default is list, not set 

9/16 : start the move to enumerate => bag
function.cl : new code for Core/enumerate
Compiler: gstat.cl -> g_stat@For, priority for lists
rebuild Claire1; rerun tests, 
   add list{x in list(1,2,3) | true } = list(1,2,3)

9/18 : claire2 week-end :)
-> 160 is the code of a " " char that happens once in a while after | on the mac
   make it invisible
-> all bu tests are OK for claire1

BIG OLD BUG: If(test = true, arg = 1, other = unknown)  -> default false is applied !
complete(o,l) -> new_writes(o,l)  returns list of p such that o.p = unknown % l
complete!(o) -> new_defaults(o,l) takes lp as a parameter

BIG NEW BUG: OID (= EID{x,O} should never be used for floats and ints) - make sure that .toEID() is used
  (a) in Kernel
  (b) in compiled code
  We need a way to check this : BadI(x,"tag") ...added before Push

9/19 : focus on bu* first 
- load(Optimize + Generate) works !
- checked ClBag.go for all methods with range any -> EID form should use r.toEID()
- good news: we get the bu* 

9/24 : start speed optimization => start with CLAIRE 1 !
- regression with Fib(30) ? 600 with stupid test on eval !409 with clean push & eval

-> regarder les notes dans le cahier précédent
call.go
self_eval@Call_method2 : looks clean
self_eval@Call

control.cl
self_eval@If

method.cl
eval_message_property : looks good

conclusion : mFib OK : as good with CLAIRE2 than CLAIRE1
Counter[0] Elapsed time: 404ms. 
Counter[0] Elapsed time: 4381ms. 
Counter[0] Elapsed time: 4689ms. 

same for mList -> checked last test 
tlt(1M)Counter[0] Elapsed time: 531ms. 
tla(10K)Counter[0] Elapsed time: 2062ms. 
tlr(200)Counter[0] Elapsed time: 525ms. 
tlw(100K)Counter[0] Elapsed time: 946ms.
last one is slightly slower

Aha mObj on second test (clr) => degradation - look at self_eval@Call+
   quite normal : we miss the Result = self.Selector.ReadEID(g0142UU)} !!!!
   Optimize/restriction!(read,list(property,any),true) fails -> return ambiguous

9/25 : mObj is now OK since open_required in restriction! (ocall.cl) is based on p
clo(1M)Counter[0] Elapsed time: 466ms. 
clr(1M)Counter[0] Elapsed time: 369ms. 
clw(1M)Counter[0] Elapsed time: 411ms. 

mSet -> exactly the same
tsta(2000) Counter[0] Elapsed time: 944ms. 
tstm(100000) Counter[0] Elapsed time: 3409ms. 

mDict -> better on all account (marginal)
tls(1000)Counter[0] Elapsed time: 908ms. 
tdt(10000)Counter[0] Elapsed time: 574ms. 
tdr(10000)Counter[0] Elapsed time: 470ms.

mSend -> exactly the same => but check with previous values

todo weekend : 
- re-run a profiling + compare with previous numbers
   -> mFib : 397 / 4400 / 4700
   -> mList : small gain :)
   -> mSend: too many MakeInteger
       (a) optimize c.MakeInt
       (b) tune compiler with g_test(quote(let r:EID := nil in (for x in list() r := eval(x), r))))
=> we get a good perf back on mSend.
=> clo is improved !
Note that debug mode is 8 times slower (but it makes sense)
=> self_eval@For on Interval : dual test of error -> for x in (1 .. 10) (eval(1), eval(2))))

1/10/2021 ====================== start claire3 ===================================

two big things to fix in claire1: 
- nb_line for multi-line C comments
    n_line() = slot for ClEnv
    add the property + compiler hook
- type!() -> type_expression
  AHA ! now that we have separated types and type_expression (extended with Patterns)
  this requires some thinking: is m.domain a list<type> or list<type_expression>

2/10/2021
(a) type!(Pattern) fixed to return any  -> range(variable) is type
    notice that glb is the extension of ^ to type_expression
    todo: add c_code for type_expression  (works for type_operator but misses Patterns)
=> compile(Generate) OK
--- break for a couple of days --------------

8/10/2021 resume on Optimize
- introduce _ before module import ? does not work for types !
  THERE IS A GO BUG : _ works for imported functions, not imported types
  AXIOM: if we use m.ClaireClass, we need an explicit "import m" at the file level
  we use types in two places: 
     - class declaration      => easy, compute the list of imported class in gosystem
       (look at all nexw classes, their super -> name -> defined)
     - variable declaration   => hard (we only extract signature automatically, the rest 
      is done with a pragma)

  Solution (a) create do_import[s:string] : set<module> in read.cl
           (b) auto declaration in define.cl 
           (c) in gosystem, when starting a file, get the list of explicit imports
           (d) later : add a slot in module  (Kernel !)
               imports(m:module) : map
               m.imports = {f1:{m1,m2}, f2:{m1,"lib"}}

  -> this fails (_ is a dumb pragma)
  better force the use (same as the _= v that the compiler generate)
  THIS WORKS PRETTY WELL : created dumb_g00 at the beginning of each generated __FILE__

  // ==================== ARGL ==================================
  We are back with an old problem : imported functions should be different from generated
  functions because they should not be prefixed by modules !
  exemple: gentenv_string
  Surprisingly, the patch was not implemented !
  (a) when importing with #<quote>f -> add a # (marker, cannot be ! because of !=)
  (b) when compiling, keep the #
  (c) when using a function name (gogen) -> if first char is # print only cdr(name)

However, this requires to boostrap (fix all existing methods with native functions)
plan
  -> produce claire1 with proper ! in front of imported through recognition in odefine.cl
     imported_function(f) -> !f (idempotent)
     imported_function?(f) : boolean
     c_princ(f) -> safe c_princ (drop first char if !)
  -> make sure that claire1, when it reads a function(nnnn), adds the !
  -> recompile claire2 : make sure that c_princ(f) cleans f
  -> added import_princ(s) which cleans the # marker for a string s

//  ================ DONE at 8:00 on Octob 10th :) ============================

// restart on Oct 14th (no battery !)
- Optimize OK !
- need to make .it exportable -> .It

// 10/16/2021
- get@list is defined with index_list, should use #index_list to avoid using Get
- C name of "C_princ" is "CPrinc"
- Generate OK

//10/17/2021
HYPOTHESIS : we do not need intermediate *.mk any more, go is doing the work :)
- gomain.cl 
   => create a name.go system file (claire1.go)
- recreate claire1 with (a) the proper module constructor + new Load function
- load gomain.cl + call system_file(Generate,"claire2")
- compile the result
- add the necessary code so that ClEnv.Params is set up in Bootstrap

// 10/19/2021
we need a consistency in module access in the Load() file
   -> if we define, assign to Module.C_variable
   -> then use it consistently
   think about iClaire which is not a package.
   iClaire was Core.C_iClaire before

// 10/23/2021
compile with claire3 (compiled compiler) !
- compile(Core)
  ARGL => bug with g_stat(self:Handle)   

10/23/2021 : claire3 files compiled !
- odefine.cl  : default is no longer unknown (look at define.cl)

10/28/2021 : restart
- claire1 from Makefile works :) + tests OK
- claire3 -> compile(Optimize) gets an endless loop in gen_meta_load
   => edit go2/ *.go code
      creer une variable dans Generate
      set it to 0 in gen_meta_load
      trace the various version of g_stat

10/29-30 :  solved the github issues ! => now the code is on git
11/1/2021:  move to compiling the tests files (perfs and bugs)

11/5/2021 : restart 
(1) mFib compile with -cm => done & good values (iso-go)!
(2) move to mList => fixed the make_go_function bug (must set status)
    - added MakeListInteger/float
    - optimize AddFast for ints
(3) mObj -> attention - si on remet les exemples inutiles => erreurs de compilation !
    - bon sur lecture/ecriture
    - x4 sur instanciation (effet des strings)
(4) mSet -> tres mauvais
    à faire à Bonnieux : ClaireSet2 façon bag
(5) mDict -> OK mais trop de test d'erreur dans tls !
    on doit pouvoir faire mieux
    -> code obthenu pour nth_table n est pas bon (error protection ?): 
(6) mSend -> 77ms not bad (much better than CLAIRE 3.5) but 
    - useless MakeInteger in solOK !!   -> 40ms when fixed.
    - boolean_I_any should not be here
    - NthPut for array should avoid errors ?

11/11/2021 : reinvent ClaireSet
  -> use same pattern as Claire3  ... not as fast for allocation (go vs claire own alloc) but OK
  -> need to implement ClaireSet (ListInteger/ListFloat/List/Any)

11/12 => write new ClBag + ClReflect

11/13 => recreate claire1 
  - done on 11/14 :)
  - add heterogeneous set tests in bu12 - OK 
  - fixed self_eval@Set in control.cl to generatr proper set<integer>
  with JITO, we get 337ms / 2000s which is thrice & twice faster than before

11/19-20: rebuild claire2, claire3 !
  -> ready to optimize the code for mDict :)
   (a)  / or mod with y = constant != 0  -> avoid the error protection
  -> ready to optimize mSend
  -> tune compiler to generate AddSetInteger (for mSet)
  -> re-run mObj and comment bad stuff out ....

11/22 : compile bu* 
- bu1 OK (but float equality is tricky ... )
- bu2 : lots of bugs !
     -> nth=(matrix,...) is macroexpanded (poor test with nth=)
     -> go does not allow evaluation without use // cf do(1,2,print(...)) removes 2
     -> 
   WE NEED TO UNDERSTAND WHAT WE DO WITH Update(...) too many strange patterns
   I need to test a multivalued add without trimmings, see if 
   => add in bu8 : multi_valued slots range object, integer, float !

- bu3 to bu7 fail badly ...
- bu8 & bu9 OK !

11/27 restart on claire3 compining
-> key for CLAIRE4 : JITO => we need to re-optimize call_methods

11/28 bu2, bu3, bu4 OK
 -> lots of bug in table compiling fixed !
 -> still todo : bu2 with safety = 5 but not 3 ! M2[1,2] := 3 -> macroexpansion M2.V produces an EID that is not seen

11/30
  A HUGE QUESTION IS : do we still need sort_abstract!(x) = any U x
  we decide do get rid of it ...
  seems to work -> bu6 OK

12/1 : work on bu7 
(a) see where c_code bugs ... 
    c_code(quote([fob(x:Store[of = X],y:X) : void -> x.value := y ]))
    method fob @ list<any>(Store[of:(any)], <ref:list(of)(ltype[0])>) is not known
    => a problem with a reference  => solved in method.cl (@)
(b) get rid of the ugly retreive_method => m is passed as the 6th args by c_code(odefine.cl)
    will need to cleanup
(c) make reference a true expression (avoid ugly let ... expansion)

12/2 : 
- fix bu7 (run code)
- bu8 : g_test(quote(tata(x = 2).z = 33))
    check that update.arg est bien g_func pour generer le code rapide
12/3 : OK from bu1 to bu10 !
  MAJOR DECISION : slot-covariance is allowed BUT we use rootSlot(s:slot) to retreive the original slot since this is the only one 
  known in go
     - this is the one used for class definition (gosystem)
     - we factor is the conversion in class constructors
     - Call_slot : g_exp and Update use rootSlot(s) to use the proper go type

12/4 bu11 (rules) ... started, to be completed (cf. notebook)
WE LEFT UOPDATE in gostat.cl UNCOMPLETED !!!!!!!!!!!!!!!!!!!!!!!!!!!!
  (1) learn to produce code for "column[1] := 2"
  (2) create a 2-dim defeasible array
  (3) check that bug3 still works (2-dim array)  => it does without store !!!!

12/10 : resume
12/11:
 - total? in odefine.cl fixed to ensure that the order is OK  (idea: create a warning or re-sort)
 - add arity 2 to demon functions in odefine.cl

12/13
 - added an underscore for complex class names: m/aClass => MAClass and m/AClass => M_AClass  : bu13 ok
 - bu14 and bu 15 OK
 - sudoku.cl (bu16) asks many questions
     (a) still a compilation bug !   Kernel.F_store_list does not produce an error/EID
     (b) compilation bug: instantiation C(x= 1,y= 2) triggers rules on x ..  (cf: the interpreter -> use update and not write)
     (c) I had to rename value into volue - > not nice
     (d) see if adding a filter would work : r1(x:Cell,y:integer)

First correction: store(l,i,x) should look for i in (1 .. length) and thus 
may produce an error   => need to retrofit to claire1 !  (function.cl)

to tests
 (a) all bu* (seems bu9, bu11 broken) - > 
 
 (b) this seems to work & bu16 works :)
    - > g_test(quote(store(C.line.counts,1,0)))   : check for errors
    - > g_test(quote(store(C.line.counts,1,0,true)))  : no errors

Second correction: instantiation should generate update(...) if if_write is present

then restore volue - > value in sudoku.cl

12/18-19 small improvements + compiler verbosity + compiler safety
 
12/20 : start compililing claire4 !
- from code in v4/meta
- improve the code to reduce the number of warnings
- add the four features
    1. Pair in Kernel + reader
       need to remove pair in function.cl
       : is already a special char
       create pair(x,y) in read.cl - > more patterns than Vardef
       - create map(lambda,l), Map (Construct)
       - map<t1,t2>(Pair*)   ... le type est obligatoire
       eval(map)
    
    2. (lvar)- >(e) in reader : 
       Attention à durcir la condition ")- >" pour ne pas réagir sur f(..) - >
    
    3. difference between x.s and s(x) - > is it Call+ ? 
       we want x.s to return an error if unknown ...
       note that under full optimization (O2) they are the same ...
       Call+ is OK, call(..) with a slot must use get vs read.
    
    4. slice for strings and lists + l[Pair] shortcut
        - > remove substring & sublist
        - > must recognize Vardef ! 

12/23 : 
- map becomes map_set, map is the usual lamda
- get rid of slice syntax because it conflicts with A[x:int] : B .....
- created claire4 - > all test but bu5 OK

12/24
- syntax for lambda is (x){e*}  !!! super elegant
- extend JITO to property with 12 restrictions !

12/26
- move all code (meta & compile) back to dropbox : clairev4.03
- implement c_code(self:Map) in odefine.cl
- create read_lambda(s:tring) in method.cl
- implement g_exp(self:lambda)

12/27: added debugger -D + moved to PC !

12/28: added test/rules
  - mFilter OK
  - mDinner OK but should work on it
  - mMonkey
  - mZebra

12/29: great version -> published on GitHub !

1/1: fix bugs
  - get @ symbol is deprecated but left in the code because value @ symbol is a slot in 3.5 !

1/2: publish 4.0.4 on GitHub
- also with blog post

// --------------------------------------- start 4.0.5 -------------------------------------------------------
 
1/16 => nth and nth= @ array are using the list EID functions hence defined in Kernel
 - also in define.cl added While to jito
 - results is q qsort which is on the same level as Python (5s vs 3.5s)

2/28 resume working on CLAIRE after a long break !
AHA ! we need a versioning strategy for go1/go2/go3/go4
   - keep them as they are : how to create claire4 (v4.0.4) from claire 3.5
   - go will be the new (current) version, using sclaire and claire (previous & current)

We need to fix the crazy compilation of not @ any ..
- should be debug-compile the compiler ?
- should we use sclaire1 + load the compiler ?
     => this is a nice investment for the future !
     create init1.cl   (init for claire1)
     this way we can trace and debug ... 

3/6 - work on optimized mode !
  compiler.optimize? = true => native global variables !
    (1) gexp.cl  / gogen.cl :  globalVar(v) produced v or v.value,  g_exp(v) must also adapt to the sort: any OR native
    (2) gstat.cl  => assign is more complex
    (3) odefine => do not produce the global variable instantiation !
    bu1 is OK

    HOWEVER: we have a HUGE regression, almost all bu do not compile anymore (bu2 to bu6) .. seems OK later
    todo 
    (a) fix with regular compiling
    (b) fix with -O

3/10 - resume working on 4.0.5 => make ctest.cl
  - bu1 et bu2 OK at 17:50
  - bu3 : normal : compile with -O but not -O5 (since there is an intensional error)
  - bu4 : same ! works with -O and -O2
  - bu5 : lambda compiling
      we known that read_lambda(s) does not return an error :)
  - bu6 : _CL_obj declared but not used
     in let x = f in x pattern with s = void
        hence the _=x is necessary if arg = var and s = void ..
     created Do! in otool.cl and I should use it more often
  - bu7 to bu12.cl work well :)
  - bu13 Kernel.PRINC("") used as value in bug13.go:90
      => changed in gosystem.cl : function_body => test if self:void and return CNULL


// --------------------------------------- start 4.0.6 -------------------------------------------------------

// TODO : avoid generating At(2-1)  => when i is a constant, perform the -1 at once !
created at_index(x:any) in goGen.cl

// 3/27 - CLAIRE week-end
(1) create a testCore (will grow progressively)  -> kernel_test.go in Kernel
   - runs OK but without option -race (check unsafe usage of pointers)
   - go test -covermode=atomic -coverprofile=coverage.out
     gives 27% coverage  => creating a test lib would make sense !

(2) fixed in ClUtil.go, Round(f) -> Floor(f)

// 4/3 : fixed let x := 1 in (y := 2)
Added a trap for error code 300 ... not the best ?
-> interpreted : Assign will NOT complain but trying to use y will (cf. self_eval @ Assign in call.cl that assumes self.var is a Variable)
-> in a method : lexical_bind (pretty.cl) will check and produce an Error 101

// 4/9 get rid of l.of == nil
- renamed makeNilList into makeBootList
- fixed by hand (Boot) the few boot lists 
- changed in otool.cl to stop (error 255) when a property is unknown and not open

// .........  June 2nd: start compiling GWDG !!! opening pandora's box .....  ..................................

problem 1:  call g_expression(X,s) where s is a type, not a class
// 6/5
 - fix g_member to return always a class
 - decide that goMethod?(m) is false if a slot co-exist with a method (same name)
 - this creates a problem with close ...
 - add princ@string
 - syntax.cl : 
       -> changed nextseq to ensure that "," implies reading a value after (vs (1,2,))
       -> changed readblock() to also checj that [foo() -> 12, ] is detected

// ******************** June 6th: create cinit.cl to run the compiler in interpreted mode ********************
- changed lexical_build into lexical_index to add final?:boolean argument

// Backlog : Error to fix !
- [foo() -> 12,]   -> is not trapped, then fails

// June 19th : publish to Git Hub
here are the steps for next  time
- make git copies all the files from the active directories (dropbox for claire, go/Kernel for go files) onto the
  root directory  (claire4 / src+meta+compile)
- gitAdd (script) adds all the files into git to be pushed
- git commit -m "your comment goes here"
- git push -f origin master

// move to CLAIRE v4.0.7 on June 20th 2022 ------------------------------------------------------------------

// resume for Xmas vacacations 2022 ! 
-  add m.resources: list of string that represent useful files (great to copy/upload module)
-  add two gitUpload and gitDownload methods  ? unclear (more like fragments)

// first step : recompile Claire4 using the new file directories

// test 2 key modules (load, compile, compare output with CLAIRE 3.5)
(a) MMS
- load (m3 vs m2 in Claire3)
- compile : error ! forward declarations do not work
BIG CHANGE : introduce NewClass(name,c,m) which checks if the class is already there (support forward def)
-> compiled(m3) works

(b) EMLA
- load (m5)
added ephemeral_object for upward_compatibility in object.cl
bug in gauge.cl => line 329 : a piece of code yields strange "enumerate primitive"
min.open was 2 in claire3.5 -> reset in 4.0
[foo(x:list<int>) -> 12] does not complain (list<X> may be parametric)

load(m5) fails .. nth error (inside code)
=> big error in nth+ (ClBag.go)  -> add the empty case nth+((),1,X)
Aha ! characters in CLAIRE4 requires their equality tests (only a small part is cashed)
=> now load(m5) works

// Dec 28th
compile(m5) fails with two problems ....
(1) soft: fractional @ integer (use a / 10) produces status = 1 (while code is safe)
    need to understand

(2) hard: 
g_test(quote(for x in LIST_LEVEL print(x)))  fails with a member_type
car pmember({list<float>(0.2, 0.3, 0.5, 0.7)}) fails !!
=> resolved (big bug in ClBag : set!(list<float>) ... used s := ... vs =

(3) Attention: slots cannot have the same name as methods in Go.
    thus the "simple" goMethod?(m) must check that no other restrictions is a slot (with intersecting domain)

(4) true compiler error with use of indirect read(p) where p = list[i]

// Jan 1st : MMS & EMLA OK + tests and perf

// when we restart : copy sclaire in a safe place and recompile everything 

// =================== to do backlog v4.0.8 (Spring 2023)==========================================

gomain.cl  -> dans mkdir, removed a direct use of / !!!!! (did not work on PC)

// New : add a reboot function in CLAIRE (useful to play with server)
// reboot() : recreate all objects (rebuild a clean state)

- ajouter la ligne de code dans le code généré  pour claire1.go
    (a) dans Load (à la fin)
    C_reboot.AddMethod(Signature(C_void.Id(), C_void.Id()), 0, MakeFunction1(E_reboot_void, "reboot_void"))
    (b) deux méthodes associées
    func reboot_void() {
      Bootstrap()
      Load()
      Reader.C_reader.Fromp = ClEnv.Cin
    }
    func E_reboot_void(s EID) EID {
      reboot_void()
      return EVOID
    }
	
// enrich the Go unit tests ! 
// create a nice command to launch tests from the makefile
list: 
   -  nth+(list<integer>(),1,1) = list<integer>(1)
   - set!(list<float>(1.2)) = {1.2}
test Equal('a','a'), Equal('a','b'), 

// before we make 4.08 a beta version, we need to test 3 more modules
- fbid (4G licence)
- SGSS (smart grid)
- SIFOA (not ready to put on github though)

// add to documentation
void is a superset of any (it is a compiler annotation: can return anything, should not be used
=> we need to explain this in the documentatation !
IT IS THE CONSEQUENCE OF ANY METHOD IS A FUNCTION THAT RETURNS A VALUE.
f() : void -> 12  is a good example.

// extensions
- use € versus % for belong
- play with spy
- imports is not used !!!
- talk about new 16 bits char in the claire documentation
- get rid of the test l.of == nil in a few weeks (once gwdg is compiled)
- dual definition of global constant should complain ? 
- add something for methods with no errors but poor range ...
     simplest is when we test can_throw / check that c_type <= range.
- doublon entre get@symbol et value@symbol -> get rid of get@symbol   
- Introduce CheckRange(type,Result,"cause")
- add the "close world assumption" for a Call in g_throw => restiction! is not empty
-  the c_substitution does not work with jito on !!!!! find out why 
      try g_test(for i in (1 .. (2 + 3)) print(i))
- fix the compilation of self_print @ string (self_print is open) => in gexp
  optimize print(integer) ?
- why do we set compiler.inline? to false, it should be true !
- when we detect a discrepency between actual throw and can_throw => 
    a. we should get rid of "good/bad lists" (unless we keep it as a debug back door)
    b. we should generate the code that is expected by the status
    c. a proper warning should be issued
- add a second-order type to check_in

=> Réfléchir à une méthode pour automatiser des tests Claire via Go
    Claire -ct mTests  -> run tests via un go test ...
    Suppose de bien comprendre le monde Go 





