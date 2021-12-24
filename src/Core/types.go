/***** CLAIRE Compilation of file /Users/ycaseau/claire/v4.0/meta/types.cl 
         [version 4.0.02 / safety 5] Friday 12-24-2021 *****/

package Core
import (_ "fmt"
	. "Kernel"
)

//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| types.cl                                                    |
//| Copyright (C) 1994 - 2021 Yves Caseau. All Rights Reserved  |
//| cf. copyright info in file object.cl: about()               |
//+-------------------------------------------------------------+
// --------------------------------------------------------------------
// This file contains the definition of the CLAIRE type system (a true lattice).
// that is used both at compile- and at run-time.
// --------------------------------------------------------------------
// ******************************************************************
// *  Table of contents                                             *
// *    Part 1: Common Set Methods                                  *
// *    Part 2: definition of the type operators                    *
// *    Part 3: Interface methods                                   *
// *    Part 4: Lattice methods                                     *
// *    Part 5: Type methods                                        *
// ******************************************************************
// *********************************************************************
// *   Part 1: Common Set Methods                                      *
// *********************************************************************
// ----------------------- useful methods ------------------------------
/* {1} The go function for: finite?(self:type) [status=0] */
func F_finite_ask_type (self *ClaireType ) *ClaireBoolean  { 
    // procedure body with s = boolean 
var Result *ClaireBoolean  
    if (C_set.Id() == self.Isa.Id()) { 
      Result = CTRUE
      }  else if (self.Isa.IsIn(C_list) == CTRUE) { 
      { var g0125 *ClaireList   = ToList(self.Id())
        _ = g0125
        { var arg_1 *ClaireAny  
          _ = arg_1
          { 
            var t *ClaireAny  
            _ = t
            arg_1= CFALSE.Id()
            var t_support *ClaireList  
            t_support = g0125
            t_len := t_support.Length()
            for i_it := 0; i_it < t_len; i_it++ { 
              t = t_support.At(i_it)
              if (ToBoolean(OBJ(F_CALL(C_finite_ask,ARGS(t.ToEID())))) != CTRUE) { 
                arg_1 = CTRUE.Id()
                break
                } else {
                
                } 
              } 
            } 
          Result = F_not_any(arg_1)
          } 
        } 
      }  else if (C_class.Id() == self.Isa.Id()) { 
      { var g0126 *ClaireClass   = ToClass(self.Id())
        { var n int  = g0126.Open
          { 
            /* Or stat: v="Result", loop=false */
            var v_or5 *ClaireBoolean  
            
            /* Or stat: try = @ any(n,open @ environment(<environment>)) with try:false, v="Result", loop=false */
            v_or5 = Equal(MakeInteger(n).Id(),MakeInteger(ClEnv.Open).Id())
            if (v_or5 == CTRUE) {Result = CTRUE
            } else { 
              /* Or stat: try = @ any(n,final @ environment(<environment>)) with try:false, v="Result", loop=false */
              v_or5 = Equal(MakeInteger(n).Id(),MakeInteger(ClEnv.Final).Id())
              if (v_or5 == CTRUE) {Result = CTRUE
              } else { 
                /* Or stat: try = @ any(n,closed()) with try:false, v="Result", loop=false */
                v_or5 = Equal(MakeInteger(n).Id(),ANY(F_CALL(C_Core_closed,ARGS(EID{ClEnv.Id(),0}))))
                if (v_or5 == CTRUE) {Result = CTRUE
                } else { 
                  /* Or stat: try ((= @ any(n,abstract @ environment(<environment>))) & (not @ any(for c:class in (subclass @ class(g0126)) (if (not @ any(finite? @ type(c))) break(true) else false)))) with try:false, v="Result", loop=false */
                  { 
                    var v_and9 *ClaireBoolean  
                    
                    v_and9 = Equal(MakeInteger(n).Id(),MakeInteger(ClEnv.ABSTRACT).Id())
                    if (v_and9 == CFALSE) {v_or5 = CFALSE
                    } else { 
                      { var arg_2 *ClaireAny  
                        _ = arg_2
                        { 
                          var c *ClaireClass  
                          _ = c
                          var c_iter *ClaireAny  
                          arg_2= CFALSE.Id()
                          var c_support *ClaireSet  
                          c_support = g0126.Subclass
                          for i_it := 0; i_it < c_support.Count; i_it++ { 
                            c_iter = c_support.At(i_it)
                            c = ToClass(c_iter)
                            if (F_finite_ask_type(ToType(c.Id())) != CTRUE) { 
                              arg_2 = CTRUE.Id()
                              break
                              } 
                            } 
                          } 
                        v_and9 = F_not_any(arg_2)
                        } 
                      if (v_and9 == CFALSE) {v_or5 = CFALSE
                      } else { 
                        v_or5 = CTRUE} 
                      } 
                    } 
                  if (v_or5 == CTRUE) {Result = CTRUE
                  } else { 
                    Result = CFALSE} 
                  } 
                } 
              } 
            } 
          } 
        } 
      } else {
      Result = CFALSE
      } 
    return Result} 
  
// The EID go function for: finite? @ type (throw: false) 
func E_finite_ask_type (self EID) EID { 
    return EID{F_finite_ask_type(ToType(OBJ(self)) ).Id(),0}} 
  
// making a set from an abstract_set  (CLAIRE 4 : bag is not longer a concrete type)
// this is a list since order matters in enumeration
/* {1} The go function for: enumerate(self:any) [status=1] */
func F_enumerate_any (self *ClaireAny ) EID { 
    var Result EID 
    if (self.Isa.IsIn(C_list) == CTRUE) { 
      { var g0128 *ClaireList   = ToList(self)
        _ = g0128
        Result = EID{g0128.Id(),0}
        } 
      }  else if (C_set.Id() == self.Isa.Id()) { 
      { var g0129 *ClaireSet   = ToSet(self)
        _ = g0129
        Result = EID{g0129.List_I().Id(),0}
        } 
      }  else if (self.Isa.IsIn(C_array) == CTRUE) { 
      { var g0130 *ClaireList   = ToArray(self)
        _ = g0130
        Result = EID{F_list_I_array(g0130).Id(),0}
        } 
      }  else if (C_class.Id() == self.Isa.Id()) { 
      { var g0131 *ClaireClass   = ToClass(self)
        _ = g0131
        { var l *ClaireList   = ToType(C_object.Id()).EmptyList()
          _ = l
          { 
            var c *ClaireClass  
            _ = c
            var c_iter *ClaireAny  
            var c_support *ClaireSet  
            c_support = g0131.Descendents
            for i_it := 0; i_it < c_support.Count; i_it++ { 
              c_iter = c_support.At(i_it)
              c = ToClass(c_iter)
              l = l.Append(c.Instances)
              } 
            } 
          Result = EID{l.Id(),0}
          } 
        } 
      }  else if (self.Isa.IsIn(C_Interval) == CTRUE) { 
      { var g0132 *ClaireInterval   = To_Interval(self)
        Result = EID{F_list_integer(g0132.Arg1,g0132.Arg2).Id(),0}
        } 
      }  else if (C_integer.Id() == self.Isa.Id()) { 
      { var g0133 int  = ToInteger(self).Value
        _ = g0133
        Result = EID{F_make_set_integer(g0133).List_I().Id(),0}
        } 
      }  else if (self.Isa.IsIn(C_collection) == CTRUE) { 
      { var g0134 *ClaireCollection   = ToCollection(self)
        _ = g0134
        { var arg_1 *ClaireAny  
          _ = arg_1
          var try_2 EID 
          /*g_try(v2:"try_2",loop:false) */
          try_2 = F_CALL(C_set_I,ARGS(EID{g0134.Id(),0}))
          /* ERROR PROTECTION INSERTED (arg_1-Result) */
          if ErrorIn(try_2) {Result = try_2
          } else {
          arg_1 = ANY(try_2)
          Result = EID{ToSet(arg_1).List_I().Id(),0}
          }
          } 
        } 
      } else {
      Result = ToException(C_general_error.Make(MakeString("[178] cannot enumerate ~S").Id(),MakeConstantList(self).Id())).Close()
      } 
    return Result} 
  
// The EID go function for: enumerate @ any (throw: true) 
func E_enumerate_any (self EID) EID { 
    return F_enumerate_any(ANY(self) )} 
  
// =type? is an operation (equality on types)
/* {1} The go function for: =type?(self:type,ens:type) [status=0] */
func F__equaltype_ask_any (self *ClaireType ,ens *ClaireType ) *ClaireBoolean  { 
    if ((self.Included(ens) == CTRUE) && 
        (ens.Included(self) == CTRUE)) {return CTRUE
    } else {return CFALSE}} 
  
// The EID go function for: =type? @ type (throw: false) 
func E__equaltype_ask_any (self EID,ens EID) EID { 
    return EID{F__equaltype_ask_any(ToType(OBJ(self)),ToType(OBJ(ens)) ).Id(),0}} 
  
// finds the sort associated to a type
/* {1} The go function for: sort!(x:type) [status=0] */
func F_sort_I_type (x *ClaireType ) *ClaireClass  { 
    // procedure body with s = class 
var Result *ClaireClass  
    if (C_class.Id() == x.Isa.Id()) { 
      { var g0136 *ClaireClass   = ToClass(x.Id())
        _ = g0136
        Result = g0136.Sort_I()
        } 
      } else {
      Result = x.Class_I().Sort_I()
      } 
    return Result} 
  
// The EID go function for: sort! @ type (throw: false) 
func E_sort_I_type (x EID) EID { 
    return EID{F_sort_I_type(ToType(OBJ(x)) ).Id(),0}} 
  
// the membership for classes
/* {1} The go function for: %(self:any,ens:class) [status=0] */
func F__Z_any1 (self *ClaireAny ,ens *ClaireClass ) *ClaireBoolean  { 
    if (self.Isa.IsIn(ens) == CTRUE) { 
      if (CTRUE == CTRUE) {return CTRUE
      } else {return CFALSE}} else {
      if (CFALSE == CTRUE) {return CTRUE
      } else {return CFALSE}} 
    } 
  
// The EID go function for: % @ list<type_expression>(any, class) (throw: false) 
func E__Z_any1 (self EID,ens EID) EID { 
    return EID{F__Z_any1(ANY(self),ToClass(OBJ(ens)) ).Id(),0}} 
  
//
// v4.0 : belong is the unique method (static call for any) for membership
// replaces belong_to + member? in claire 3 => works on everything, collections and integer as well :)
// see belong_exp in gexp.cl to see how it is used + open-conding patterns
// note that belong may create an error => heavier => optimize with %t when possible 
/* {1} The go function for: belong(x:any,y:any) [status=1] */
func F_BELONG (x *ClaireAny ,y *ClaireAny ) EID { 
    var Result EID 
    if (C_class.Id() == y.Isa.Id()) { 
      { var g0138 *ClaireClass   = ToClass(y)
        _ = g0138
        Result = EID{F__Z_any1(x,g0138).Id(),0}
        } 
      }  else if (y.Isa.IsIn(C_list) == CTRUE) { 
      { var g0139 *ClaireList   = ToList(y)
        _ = g0139
        Result = EID{g0139.Contain_ask(x).Id(),0}
        } 
      }  else if (C_set.Id() == y.Isa.Id()) { 
      { var g0140 *ClaireSet   = ToSet(y)
        _ = g0140
        Result = EID{ToType(g0140.Id()).Contains(x).Id(),0}
        } 
      }  else if (y.Isa.IsIn(C_array) == CTRUE) { 
      { var g0141 *ClaireList   = ToArray(y)
        _ = g0141
        Result = EID{ToList(g0141.Id()).Contain_ask(x).Id(),0}
        } 
      }  else if (C_tuple.Id() == y.Isa.Id()) { 
      { var g0142 *ClaireTuple   = ToTuple(y)
        { 
          var v_and4 *ClaireBoolean  
          
          v_and4 = Equal(C_tuple.Id(),x.Isa.Id())
          if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
          } else { 
            v_and4 = Equal(ANY(F_CALL(C_length,ARGS(x.ToEID()))),MakeInteger(g0142.Length()).Id())
            if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
            } else { 
              var try_1 EID 
              /*g_try(v2:"try_1",loop:false) */
              { var arg_2 *ClaireAny  
                _ = arg_2
                var try_3 EID 
                /*g_try(v2:"try_3",loop:false) */
                { var i int  = 1
                  { var g0143 int  = INT(F_CALL(C_length,ARGS(x.ToEID())))
                    _ = g0143
                    try_3= EID{CFALSE.Id(),0}
                    for (i <= g0143) { 
                      /* While stat, v:"try_3" loop:false */
                      var loop_4 EID 
                      _ = loop_4
                      { 
                      /*g_try(v2:"loop_4",loop:tuple("try_3", EID)) */
                      var g0150I *ClaireBoolean  
                      var try_5 EID 
                      /*g_try(v2:"try_5",loop:false) */
                      { var arg_6 *ClaireBoolean  
                        _ = arg_6
                        var try_7 EID 
                        /*g_try(v2:"try_7",loop:false) */
                        { var arg_8 *ClaireAny  
                          _ = arg_8
                          var try_9 EID 
                          /*g_try(v2:"try_9",loop:false) */
                          try_9 = F_CALL(C_nth,ARGS(x.ToEID(),EID{C__INT,IVAL(i)}))
                          /* ERROR PROTECTION INSERTED (arg_8-try_7) */
                          if ErrorIn(try_9) {try_7 = try_9
                          } else {
                          arg_8 = ANY(try_9)
                          try_7 = F_BELONG(arg_8,ToList(g0142.Id()).At(i-1))
                          }
                          } 
                        /* ERROR PROTECTION INSERTED (arg_6-try_5) */
                        if ErrorIn(try_7) {try_5 = try_7
                        } else {
                        arg_6 = ToBoolean(OBJ(try_7))
                        try_5 = EID{arg_6.Not.Id(),0}
                        }
                        } 
                      /* ERROR PROTECTION INSERTED (g0150I-loop_4) */
                      if ErrorIn(try_5) {loop_4 = try_5
                      } else {
                      g0150I = ToBoolean(OBJ(try_5))
                      if (g0150I == CTRUE) { 
                        try_3 = EID{CTRUE.Id(),0}
                        break
                        } else {
                        loop_4 = EID{CFALSE.Id(),0}
                        } 
                      }
                      /* ERROR PROTECTION INSERTED (loop_4-loop_4) */
                      if ErrorIn(loop_4) {try_3 = loop_4
                      break
                      } else {
                      i = (i+1)
                      }
                      /* try?:false, v2:"v_while10" loop will be:tuple("try_3", EID) */
                      } 
                    }
                    } 
                  } 
                /* ERROR PROTECTION INSERTED (arg_2-try_1) */
                if ErrorIn(try_3) {try_1 = try_3
                } else {
                arg_2 = ANY(try_3)
                try_1 = EID{F_not_any(arg_2).Id(),0}
                }
                } 
              /* ERROR PROTECTION INSERTED (v_and4-Result) */
              if ErrorIn(try_1) {Result = try_1
              } else {
              v_and4 = ToBoolean(OBJ(try_1))
              if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
              } else { 
                Result = EID{CTRUE.Id(),0}} 
              } 
            } 
          }
          } 
        } 
      }  else if (y.Isa.IsIn(C_type_operator) == CTRUE) { 
      { var g0144 *ClaireTypeOperator   = ToTypeOperator(y)
        _ = g0144
        Result = EID{ToType(g0144.Id()).Contains(x).Id(),0}
        } 
      }  else if (C_integer.Id() == y.Isa.Id()) { 
      { var g0145 int  = ToInteger(y).Value
        _ = g0145
        if (C_integer.Id() == x.Isa.Id()) { 
          { var g0146 int  = ToInteger(x).Value
            _ = g0146
            Result = EID{BitVectorContains(g0145,g0146).Id(),0}
            } 
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        } 
      } else {
      { var start int  = ClEnv.Index
        ClEnv.Push(x.ToEID())
        ClEnv.Push(y.ToEID())
        { var m *ClaireObject   = F_find_which_property(ToProperty(C__Z.Id()),start,x.Isa)
          var g0151I *ClaireBoolean  
          if (C_method.Id() == m.Isa.Id()) { 
            { var g0149 *ClaireMethod   = ToMethod(m.Id())
              g0151I = MakeBoolean((g0149.Domain.Length() == 2) && (g0149.Domain.ValuesO()[2-1] != C_any.Id()))
              } 
            } else {
            g0151I = CFALSE
            } 
          if (g0151I == CTRUE) { 
            Result = F_eval_message_property(ToProperty(C__Z.Id()),m,start,CTRUE)
            } else {
            Result = ToException(C_general_error.Make(MakeString("[179] (~S % ~S): not implemented!").Id(),MakeConstantList(x,y).Id())).Close()
            } 
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: belong @ any (throw: true) 
func E_BELONG (x EID,y EID) EID { 
    return F_BELONG(ANY(x),ANY(y) )} 
  
// x % y is a short cut 
// CLAIRE4 : cannot be a macro (too early)
/* {1} The go function for: %(x:any,y:any) [status=1] */
func F_belong_to (x *ClaireAny ,y *ClaireAny ) EID { 
    var Result EID 
    Result = F_BELONG(x,y)
    return Result} 
  
// The EID go function for: % @ list<type_expression>(any, any) (throw: true) 
func E_belong_to (x EID,y EID) EID { 
    return F_belong_to(ANY(x),ANY(y) )} 
  
// ****************************************************************
// *         Part 2: definition of the type operators             *
// ****************************************************************
// in CLAIRE4, types are defined in the Kernel go module
// type_operator <: type()
// union of two types ---------------------------------------------
// Disjonctive Union Axiom (DU): Each union (A U B) is stricly disjunctive:
//       (1) A ^B = 0
//       (2) x < A U B <=> x < A or x < B
// Producing disjunction union is a form of normalization (the previous notion
// of diustributivity was a lousy bug)
// DU Axiom is necessary to make <= and ^ easier to define
// This is achieved in the U method
/* {1} The go function for: self_print(self:Union) [status=1] */
func F_self_print_Union_Core (self *ClaireUnion ) EID { 
    var Result EID 
    PRINC("(")
    /*g_try(v2:"Result",loop:true) */
    Result = F_print_any(self.T1.Id())
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" U ")
    /*g_try(v2:"Result",loop:true) */
    Result = F_print_any(self.T2.Id())
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(")")
    Result = EVOID
    }}
    return Result} 
  
// The EID go function for: self_print @ Union (throw: true) 
func E_self_print_Union_Core (self EID) EID { 
    return F_self_print_Union_Core(To_Union(OBJ(self)) )} 
  
/* {1} The go function for: finite?(self:Union) [status=0] */
func F_finite_ask_Union (self *ClaireUnion ) *ClaireBoolean  { 
    if ((ToBoolean(OBJ(F_CALL(C_finite_ask,ARGS(EID{self.T1.Id(),0})))) == CTRUE) && 
        (ToBoolean(OBJ(F_CALL(C_finite_ask,ARGS(EID{self.T2.Id(),0})))) == CTRUE)) {return CTRUE
    } else {return CFALSE}} 
  
// The EID go function for: finite? @ Union (throw: false) 
func E_finite_ask_Union (self EID) EID { 
    return EID{F_finite_ask_Union(To_Union(OBJ(self)) ).Id(),0}} 
  
// Intervals of integers ----------
/* {1} The go function for: self_print(self:Interval) [status=1] */
func F_self_print_Interval_Core (self *ClaireInterval ) EID { 
    var Result EID 
    PRINC("(")
    /*g_try(v2:"Result",loop:true) */
    Result = F_print_any(MakeInteger(self.Arg1).Id())
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" .. ")
    /*g_try(v2:"Result",loop:true) */
    Result = F_print_any(MakeInteger(self.Arg2).Id())
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(")")
    Result = EVOID
    }}
    return Result} 
  
// The EID go function for: self_print @ Interval (throw: true) 
func E_self_print_Interval_Core (self EID) EID { 
    return F_self_print_Interval_Core(To_Interval(OBJ(self)) )} 
  
/* {1} The go function for: finite?(self:Interval) [status=0] */
func F_finite_ask_Interval (self *ClaireInterval ) *ClaireBoolean  { 
    if (CTRUE == CTRUE) {return CTRUE
    } else {return CFALSE}} 
  
// The EID go function for: finite? @ Interval (throw: false) 
func E_finite_ask_Interval (self EID) EID { 
    return EID{F_finite_ask_Interval(To_Interval(OBJ(self)) ).Id(),0}} 
  
// true constructor
/* {1} The go function for: --(x:integer,y:integer) [status=1] */
func F__dash_dash_integer (x int,y int) EID { 
    var Result EID 
    if (x <= y) { 
      Result = EID{F__dot_dot_integer(x,y).Id(),0}
      } else {
      Result = ToException(C_general_error.Make(MakeString("[182] the interval (~S -- ~S) is empty").Id(),MakeConstantList(MakeInteger(x).Id(),MakeInteger(y).Id()).Id())).Close()
      } 
    return Result} 
  
// The EID go function for: -- @ integer (throw: true) 
func E__dash_dash_integer (x EID,y EID) EID { 
    return F__dash_dash_integer(INT(x),INT(y) )} 
  
// Parameterized class. -------------------------------------------
/* {1} The go function for: self_print(self:Param) [status=1] */
func F_self_print_Param_Core (self *ClaireParam ) EID { 
    var Result EID 
    if ((self.Params.Length() == 1) && 
        ((self.Params.At(1-1) == C_of.Id()) && 
          (C_set.Id() == self.Args.At(1-1).Isa.Id()))) { 
      /*g_try(v2:"Result",loop:true) */
      Result = F_print_any(self.Arg.Id())
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC("<")
      /*g_try(v2:"Result",loop:true) */
      { var arg_1 *ClaireAny  
        _ = arg_1
        var try_2 EID 
        /*g_try(v2:"try_2",loop:false) */
        try_2 = F_the_type(ToType(self.Args.At(1-1)))
        /* ERROR PROTECTION INSERTED (arg_1-Result) */
        if ErrorIn(try_2) {Result = try_2
        } else {
        arg_1 = ANY(try_2)
        Result = F_CALL(C_print,ARGS(arg_1.ToEID()))
        }
        } 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(">")
      Result = EVOID
      }}
      } else {
      /*g_try(v2:"Result",loop:true) */
      Result = F_print_any(self.Arg.Id())
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC("[")
      /*g_try(v2:"Result",loop:true) */
      { var i int  = 1
        { var g0152 int  = self.Args.Length()
          _ = g0152
          Result= EID{CFALSE.Id(),0}
          for (i <= g0152) { 
            /* While stat, v:"Result" loop:true */
            var loop_3 EID 
            _ = loop_3
            { 
            /*g_try(v2:"loop_3",loop:tuple("Result", EID)) */
            if (i != 1) { 
              PRINC(", ")
              } 
            /*g_try(v2:"loop_3",loop:tuple("Result", EID)) */
            /*g_try(v2:"loop_3",loop:tuple("Result", EID)) */
            loop_3 = F_CALL(C_print,ARGS(self.Params.At(i-1).ToEID()))
            /* ERROR PROTECTION INSERTED (loop_3-loop_3) */
            if ErrorIn(loop_3) {Result = loop_3
            break
            } else {
            PRINC(":(")
            /*g_try(v2:"loop_3",loop:tuple("Result", EID)) */
            loop_3 = F_CALL(C_print,ARGS(self.Args.At(i-1).ToEID()))
            /* ERROR PROTECTION INSERTED (loop_3-loop_3) */
            if ErrorIn(loop_3) {Result = loop_3
            break
            } else {
            PRINC(")")
            loop_3 = EVOID
            }}
            /* ERROR PROTECTION INSERTED (loop_3-loop_3) */
            if ErrorIn(loop_3) {Result = loop_3
            break
            } else {
            }
            /* ERROR PROTECTION INSERTED (loop_3-loop_3) */
            if ErrorIn(loop_3) {Result = loop_3
            break
            } else {
            i = (i+1)
            }
            /* try?:false, v2:"v_while5" loop will be:tuple("Result", EID) */
            } 
          }
          } 
        } 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC("]")
      Result = EVOID
      }}
      } 
    return Result} 
  
// The EID go function for: self_print @ Param (throw: true) 
func E_self_print_Param_Core (self EID) EID { 
    return F_self_print_Param_Core(To_Param(OBJ(self)) )} 
  
/* {1} The go function for: finite?(self:Param) [status=0] */
func F_finite_ask_Param (self *ClaireParam ) *ClaireBoolean  { 
    if (F_finite_ask_type(ToType(self.Arg.Id())) == CTRUE) {return CTRUE
    } else {return CFALSE}} 
  
// The EID go function for: finite? @ Param (throw: false) 
func E_finite_ask_Param (self EID) EID { 
    return EID{F_finite_ask_Param(To_Param(OBJ(self)) ).Id(),0}} 
  
// subtype[X] ----------------------------------------------
// subtype[X] = {u in type | u <= t}
// for closure purposes, we add an arg Y -> Y inter st[X]
// Y can be any type class, but we forbid parametrisation on such classes !
// thus we can ensure that Y is a class
/* {1} The go function for: self_print(self:subtype) [status=1] */
func F_self_print_subtype_Core (self *ClaireSubtype ) EID { 
    var Result EID 
    if (self.Arg.Id() == C_type.Id()) { 
      PRINC("subtype[")
      /*g_try(v2:"Result",loop:true) */
      Result = F_print_any(self.T1.Id())
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC("]")
      Result = EVOID
      }
      } else {
      /*g_try(v2:"Result",loop:true) */
      Result = F_print_any(self.Arg.Id())
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC("[")
      /*g_try(v2:"Result",loop:true) */
      Result = F_print_any(self.T1.Id())
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC("]")
      Result = EVOID
      }}
      } 
    return Result} 
  
// The EID go function for: self_print @ subtype (throw: true) 
func E_self_print_subtype_Core (self EID) EID { 
    return F_self_print_subtype_Core(ToSubtype(OBJ(self)) )} 
  
// v3.2
/* {1} The go function for: finite?(self:subtype) [status=0] */
func F_finite_ask_subtype (self *ClaireSubtype ) *ClaireBoolean  { 
    if ((self.Arg.Id() == C_set.Id()) && 
        (ToBoolean(OBJ(F_CALL(C_finite_ask,ARGS(EID{self.T1.Id(),0})))) == CTRUE)) {return CTRUE
    } else {return CFALSE}} 
  
// The EID go function for: finite? @ subtype (throw: false) 
func E_finite_ask_subtype (self EID) EID { 
    return EID{F_finite_ask_subtype(ToSubtype(OBJ(self)) ).Id(),0}} 
  
// creates a subtype, with some normalization
// v3.2 list[t] -> subtype 
// v4.0 => no error
/* {1} The go function for: nth(self:class,x:type) [status=0] */
func F_nth_class1 (self *ClaireClass ,x *ClaireType ) *ClaireType  { 
    // procedure body with s = type 
var Result *ClaireType  
    if ((self.Id() == C_set.Id()) || 
        (self.Id() == C_list.Id())) { 
      { var _CL_obj *ClaireSubtype   = ToSubtype(new(ClaireSubtype).Is(C_subtype))
        _CL_obj.Arg = self
        /*class->class*/_CL_obj.T1 = x
        /*type->type*/Result = ToType(_CL_obj.Id())
        } 
      }  else if (self.IsIn(C_type) != CTRUE) { 
      Result = ToType(CEMPTY.Id())
      } else {
      { var _CL_obj *ClaireSubtype   = ToSubtype(new(ClaireSubtype).Is(C_subtype))
        { 
          var va_arg1 *ClaireSubtype  
          var va_arg2 *ClaireClass  
          va_arg1 = _CL_obj
          if (self.Id() == C_subtype.Id()) { 
            va_arg2 = C_type
            } else {
            va_arg2 = self
            } 
          va_arg1.Arg = va_arg2
          /*class->class*/} 
        _CL_obj.T1 = x
        /*type->type*/Result = ToType(_CL_obj.Id())
        } 
      } 
    return Result} 
  
// The EID go function for: nth @ list<type_expression>(class, type) (throw: false) 
func E_nth_class1 (self EID,x EID) EID { 
    return EID{F_nth_class1(ToClass(OBJ(self)),ToType(OBJ(x)) ).Id(),0}} 
  
// create a Param with a list of parameters (constant properties) l1 and a list
// of types l2
// v4.0 => no error
/* {1} The go function for: nth(self:class,l1:list,l2:list) [status=0] */
func F_nth_class2 (self *ClaireClass ,l1 *ClaireList ,l2 *ClaireList ) *ClaireType  { 
    // procedure body with s = type 
var Result *ClaireType  
    if (((self.Id() == C_list.Id()) || 
          (self.Id() == C_set.Id())) && 
        (l2.At(1-1).Isa.IsIn(C_subtype) == CTRUE)) { 
      Result = F_nth_class1(self,ToSubtype(l2.At(1-1)).T1)
      }  else if (((self.Id() == C_list.Id()) || 
          (self.Id() == C_set.Id())) && 
        (l1.At(1-1) != C_of.Id())) { 
      Result = ToType(CEMPTY.Id())
      } else {
      { var _CL_obj *ClaireParam   = To_Param(new(ClaireParam).Is(C_Param))
        _CL_obj.Arg = self
        /*class->class*/_CL_obj.Params = l1
        /*list->list*/_CL_obj.Args = l2
        /*list->list*/Result = ToType(_CL_obj.Id())
        } 
      } 
    return Result} 
  
// The EID go function for: nth @ list<type_expression>(class, list, list) (throw: false) 
func E_nth_class2 (self EID,l1 EID,l2 EID) EID { 
    return EID{F_nth_class2(ToClass(OBJ(self)),ToList(OBJ(l1)),ToList(OBJ(l2)) ).Id(),0}} 
  
// create a Param of the stack[X] kind
/* {1} The go function for: param!(self:class,tx:type) [status=0] */
func F_param_I_class (self *ClaireClass ,tx *ClaireType ) *ClaireType  { 
    // procedure body with s = type 
var Result *ClaireType  
    { var _CL_obj *ClaireParam   = To_Param(new(ClaireParam).Is(C_Param))
      _CL_obj.Arg = self
      /*class->class*/_CL_obj.Params = MakeConstantList(C_of.Id())
      /*list->list*/_CL_obj.Args = MakeConstantList(MakeConstantSet(tx.Id()).Id())
      /*list->list*/Result = ToType(_CL_obj.Id())
      } 
    return Result} 
  
// The EID go function for: param! @ class (throw: false) 
func E_param_I_class (self EID,tx EID) EID { 
    return EID{F_param_I_class(ToClass(OBJ(self)),ToType(OBJ(tx)) ).Id(),0}} 
  
// create the t[] param
/* {1} The go function for: nth(self:type) [status=0] */
func F_nth_type (self *ClaireType ) *ClaireType  { 
    // procedure body with s = type 
var Result *ClaireType  
    { var _CL_obj *ClaireParam   = To_Param(new(ClaireParam).Is(C_Param))
      _CL_obj.Arg = C_array
      /*class->class*/_CL_obj.Params = MakeConstantList(C_of.Id())
      /*list->list*/_CL_obj.Args = MakeConstantList(MakeConstantSet(self.Id()).Id())
      /*list->list*/Result = ToType(_CL_obj.Id())
      } 
    return Result} 
  
// The EID go function for: nth @ type (throw: false) 
func E_nth_type (self EID) EID { 
    return EID{F_nth_type(ToType(OBJ(self)) ).Id(),0}} 
  
// tuple are types
/* {1} The go function for: finite?(self:tuple) [status=0] */
func F_finite_ask_tuple (self *ClaireTuple ) *ClaireBoolean  { 
    // procedure body with s = boolean 
var Result *ClaireBoolean  
    { var arg_1 *ClaireAny  
      _ = arg_1
      { 
        var x *ClaireAny  
        _ = x
        arg_1= CFALSE.Id()
        var x_support *ClaireList  
        x_support = ToList(self.Id())
        x_len := x_support.Length()
        for i_it := 0; i_it < x_len; i_it++ { 
          x = x_support.At(i_it)
          if (ToBoolean(OBJ(F_CALL(C_finite_ask,ARGS(x.ToEID())))) != CTRUE) { 
            arg_1 = CTRUE.Id()
            break
            } 
          } 
        } 
      Result = F_not_any(arg_1)
      } 
    return Result} 
  
// The EID go function for: finite? @ tuple (throw: false) 
func E_finite_ask_tuple (self EID) EID { 
    return EID{F_finite_ask_tuple(ToTuple(OBJ(self)) ).Id(),0}} 
  
// reference to a previous variable, not a type but a pattern -------
// index is the position of the stack of the referred type
// args is a list representing the path (a sequence of properties (parameters))
// a property is applied to the referred type
// if arg = true, the reference is the singleton containing the ref. value
// arg is set to true when we copy a reference in define.cl (unclear why)
/* {1} The go function for: self_print(self:Reference) [status=1] */
func F_self_print_Reference_Core (self *ClaireReference ) EID { 
    var Result EID 
    PRINC("<ref:")
    /*g_try(v2:"Result",loop:true) */
    Result = F_print_any(self.Args.Id())
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("(ltype[")
    F_princ_integer(self.Index)
    PRINC("])>")
    Result = EVOID
    }
    return Result} 
  
// The EID go function for: self_print @ Reference (throw: true) 
func E_self_print_Reference_Core (self EID) EID { 
    return F_self_print_Reference_Core(To_Reference(OBJ(self)) )} 
  
/* {1} The go function for: get(self:Reference,y:any) [status=0] */
func F_get_Reference (self *ClaireReference ,y *ClaireAny ) *ClaireAny  { 
    // procedure body with s = any 
var Result *ClaireAny  
    { var l *ClaireList   = self.Args
      _ = l
      { var i int  = 1
        _ = i
        { var g0153 int  = l.Length()
          _ = g0153
          for (i <= g0153) { 
            /* While stat, v:"Result" loop:false */
            y = ANY(F_funcall_property(ToProperty(l.At(i-1)),y))
            i = (i+1)
            /* try?:false, v2:"v_while5" loop will be:tuple("Result", void) */
            } 
          } 
        } 
      Result = y
      } 
    return Result} 
  
// The EID go function for: get @ Reference (throw: false) 
func E_get_Reference (self EID,y EID) EID { 
    return F_get_Reference(To_Reference(OBJ(self)),ANY(y) ).ToEID()} 
  
// we need a constructor
/* {1} The go function for: Reference!(l:list,n:integer) [status=0] */
func F_Reference_I_list (l *ClaireList ,n int) *ClaireReference  { 
    // procedure body with s = Reference 
var Result *ClaireReference  
    { var _CL_obj *ClaireReference   = To_Reference(new(ClaireReference).Is(C_Reference))
      _CL_obj.Args = l
      /*list->list*/_CL_obj.Index = n
      /*integer->integer*/Result = _CL_obj
      } 
    return Result} 
  
// The EID go function for: Reference! @ list (throw: false) 
func E_Reference_I_list (l EID,n EID) EID { 
    return EID{F_Reference_I_list(ToList(OBJ(l)),INT(n) ).Id(),0}} 
  
// apply a reference to a type (l is args(self), passed for disambiguation)
/* {1} The go function for: @(self:Reference,l:list,y:any) [status=0] */
func F__at_Reference (self *ClaireReference ,l *ClaireList ,y *ClaireAny ) *ClaireAny  { 
    
    { var i int  = 1
      _ = i
      { var g0154 int  = l.Length()
        _ = g0154
        for (i <= g0154) { 
          /* While stat, v:"Unused" loop:false */
          y = ToType(y).At(ToProperty(l.At(i-1))).Id()
          i = (i+1)
          /* try?:false, v2:"v_while4" loop will be:tuple("Unused", void) */
          } 
        } 
      } 
    return  y
    } 
  
// The EID go function for: @ @ Reference (throw: false) 
func E__at_Reference (self EID,l EID,y EID) EID { 
    return F__at_Reference(To_Reference(OBJ(self)),ToList(OBJ(l)),ANY(y) ).ToEID()} 
  
// type to set coercion  -------------------------------------------------
// new in v3.0.5 = use an interface method for type enumeration
// the default strategy is extensible: we look if there exists
// a proper definition that could be interpreted !
/* {1} The go function for: set!(x:collection) [status=1] */
func F_set_I_collection (x *ClaireCollection ) EID { 
    var Result EID 
    { var m *ClaireAny   = F__at_property1(C_set_I,x.Isa).Id()
      if (F_domain_I_restriction(ToRestriction(m)).Id() != C_collection.Id()) { 
        Result = F_CALL(C_funcall,ARGS(m.ToEID(),EID{x.Id(),0}))
        } else {
        Result = ToException(C_general_error.Make(MakeString("[178] cannot enumerate ~S").Id(),MakeConstantList(x.Id()).Id())).Close()
        } 
      } 
    return Result} 
  
// The EID go function for: set! @ collection (throw: true) 
func E_set_I_collection (x EID) EID { 
    return F_set_I_collection(ToCollection(OBJ(x)) )} 
  
/* {1} The go function for: size(x:collection) [status=1] */
func F_size_collection (x *ClaireCollection ) EID { 
    var Result EID 
    { var m *ClaireAny   = F__at_property1(C_size,x.Isa).Id()
      if (F_domain_I_restriction(ToRestriction(m)).Id() != C_collection.Id()) { 
        Result = F_CALL(C_funcall,ARGS(m.ToEID(),EID{x.Id(),0}))
        } else {
        { var arg_1 *ClaireAny  
          _ = arg_1
          var try_2 EID 
          /*g_try(v2:"try_2",loop:false) */
          try_2 = F_CALL(C_set_I,ARGS(EID{x.Id(),0}))
          /* ERROR PROTECTION INSERTED (arg_1-Result) */
          if ErrorIn(try_2) {Result = try_2
          } else {
          arg_1 = ANY(try_2)
          Result = EID{C__INT,IVAL(ToSet(arg_1).Size())}
          }
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: size @ collection (throw: true) 
func E_size_collection (x EID) EID { 
    return F_size_collection(ToCollection(OBJ(x)) )} 
  
// v3.2.34  -> makes the API simpler
// (interface(size))
// set is needed for recursive def
/* {1} The go function for: set!(x:set) [status=0] */
func F_set_I_set (x *ClaireSet ) *ClaireSet  { 
    return  x
    } 
  
// The EID go function for: set! @ set (throw: false) 
func E_set_I_set (x EID) EID { 
    return EID{F_set_I_set(ToSet(OBJ(x)) ).Id(),0}} 
  
// set is needed for recursive def
/* {1} The go function for: size(x:list) [status=0] */
func F_size_list2_Core (x *ClaireList ) int { 
    return  x.Set_I().Size()
    } 
  
// The EID go function for: size @ list (throw: false) 
func E_size_list2_Core (x EID) EID { 
    return EID{C__INT,IVAL(F_size_list2_Core(ToList(OBJ(x)) ))}} 
  
// class  -> return a read-only list  (v3.2)
/* {1} The go function for: set!(x:class) [status=1] */
func F_set_I_class (x *ClaireClass ) EID { 
    var Result EID 
    { var rep *ClaireList   = ToType(CEMPTY.Id()).EmptyList()
      _ = rep
      /*g_try(v2:"Result",loop:true) */
      { 
        var c *ClaireClass  
        _ = c
        var c_iter *ClaireAny  
        Result= EID{CFALSE.Id(),0}
        var c_support *ClaireSet  
        c_support = x.Descendents
        for i_it := 0; i_it < c_support.Count; i_it++ { 
          c_iter = c_support.At(i_it)
          c = ToClass(c_iter)
          var loop_1 EID 
          _ = loop_1
          /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
          if ((c.IsIn(C_primitive) == CTRUE) && 
              (c.Id() != C_boolean.Id())) { 
            loop_1 = ToException(C_general_error.Make(MakeString("[178] cannot enumerate ~S").Id(),MakeConstantList(c.Id()).Id())).Close()
            } else {
            rep = rep.Append(c.Instances)
            loop_1 = EID{rep.Id(),0}
            } 
          /* ERROR PROTECTION INSERTED (loop_1-Result) */
          if ErrorIn(loop_1) {Result = loop_1
          break
          } else {
          }
          } 
        } 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = EID{rep.Set_I().Id(),0}
      }
      } 
    return Result} 
  
// The EID go function for: set! @ class (throw: true) 
func E_set_I_class (x EID) EID { 
    return F_set_I_class(ToClass(OBJ(x)) )} 
  
/* {1} The go function for: size(self:class) [status=0] */
func F_size_class (self *ClaireClass ) int { 
    // procedure body with s = integer 
var Result int 
    { var n int  = 0
      _ = n
      { 
        var x *ClaireClass  
        _ = x
        var x_iter *ClaireAny  
        var x_support *ClaireSet  
        x_support = self.Descendents
        for i_it := 0; i_it < x_support.Count; i_it++ { 
          x_iter = x_support.At(i_it)
          x = ToClass(x_iter)
          n = (n+x.Instances.Length())
          } 
        } 
      Result = n
      } 
    return Result} 
  
// The EID go function for: size @ class (throw: false) 
func E_size_class (self EID) EID { 
    return EID{C__INT,IVAL(F_size_class(ToClass(OBJ(self)) ))}} 
  
// Union
/* {1} The go function for: set!(x:Union) [status=1] */
func F_set_I_Union (x *ClaireUnion ) EID { 
    var Result EID 
    { var arg_1 *ClaireAny  
      _ = arg_1
      var try_3 EID 
      /*g_try(v2:"try_3",loop:false) */
      try_3 = F_CALL(C_set_I,ARGS(EID{x.T1.Id(),0}))
      /* ERROR PROTECTION INSERTED (arg_1-Result) */
      if ErrorIn(try_3) {Result = try_3
      } else {
      arg_1 = ANY(try_3)
      { var arg_2 *ClaireAny  
        _ = arg_2
        var try_4 EID 
        /*g_try(v2:"try_4",loop:false) */
        try_4 = F_CALL(C_set_I,ARGS(EID{x.T2.Id(),0}))
        /* ERROR PROTECTION INSERTED (arg_2-Result) */
        if ErrorIn(try_4) {Result = try_4
        } else {
        arg_2 = ANY(try_4)
        Result = EID{F_append_set(ToSet(arg_1),ToSet(arg_2)).Id(),0}
        }
        } 
      }
      } 
    return Result} 
  
// The EID go function for: set! @ Union (throw: true) 
func E_set_I_Union (x EID) EID { 
    return F_set_I_Union(To_Union(OBJ(x)) )} 
  
/* {1} The go function for: size(x:Union) [status=1] */
func F_size_Union (x *ClaireUnion ) EID { 
    var Result EID 
    if ((x.T1.Isa.IsIn(C_Interval) == CTRUE) || 
        (C_set.Id() == x.T1.Isa.Id())) { 
      { var arg_1 *ClaireAny  
        _ = arg_1
        var try_3 EID 
        /*g_try(v2:"try_3",loop:false) */
        try_3 = F_CALL(C_size,ARGS(EID{x.T1.Id(),0}))
        /* ERROR PROTECTION INSERTED (arg_1-Result) */
        if ErrorIn(try_3) {Result = try_3
        } else {
        arg_1 = ANY(try_3)
        { var arg_2 *ClaireAny  
          _ = arg_2
          var try_4 EID 
          /*g_try(v2:"try_4",loop:false) */
          try_4 = F_CALL(C_size,ARGS(EID{x.T2.Id(),0}))
          /* ERROR PROTECTION INSERTED (arg_2-Result) */
          if ErrorIn(try_4) {Result = try_4
          } else {
          arg_2 = ANY(try_4)
          Result = EID{C__INT,IVAL(F__plus_integer(ToInteger(arg_1).Value,ToInteger(arg_2).Value))}
          }
          } 
        }
        } 
      } else {
      { var arg_5 *ClaireSet  
        _ = arg_5
        var try_6 EID 
        /*g_try(v2:"try_6",loop:false) */
        try_6 = F_set_I_Union(x)
        /* ERROR PROTECTION INSERTED (arg_5-Result) */
        if ErrorIn(try_6) {Result = try_6
        } else {
        arg_5 = ToSet(OBJ(try_6))
        Result = EID{C__INT,IVAL(arg_5.Size())}
        }
        } 
      } 
    return Result} 
  
// The EID go function for: size @ Union (throw: true) 
func E_size_Union (x EID) EID { 
    return F_size_Union(To_Union(OBJ(x)) )} 
  
// interval
/* {1} The go function for: set!(x:Interval) [status=0] */
func F_set_I_Interval (x *ClaireInterval ) *ClaireSet  { 
    return  ToSet(F_sequence_integer(x.Arg1,x.Arg2).Id())
    } 
  
// The EID go function for: set! @ Interval (throw: false) 
func E_set_I_Interval (x EID) EID { 
    return EID{F_set_I_Interval(To_Interval(OBJ(x)) ).Id(),0}} 
  
/* {1} The go function for: size(self:Interval) [status=0] */
func F_size_Interval (self *ClaireInterval ) int { 
    return  ((self.Arg2+1)-self.Arg1)
    } 
  
// The EID go function for: size @ Interval (throw: false) 
func E_size_Interval (self EID) EID { 
    return EID{C__INT,IVAL(F_size_Interval(To_Interval(OBJ(self)) ))}} 
  
// param
/* {1} The go function for: set!(x:Param) [status=1] */
func F_set_I_Param (x *ClaireParam ) EID { 
    var Result EID 
    { var y_in *ClaireSet  
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      try_1 = F_set_I_class(x.Arg)
      /* ERROR PROTECTION INSERTED (y_in-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      y_in = ToSet(OBJ(try_1))
      { var y_out *ClaireSet   = y_in.Empty()
        { 
          var y *ClaireAny  
          _ = y
          var y_support *ClaireSet  
          y_support = y_in
          for i_it := 0; i_it < y_support.Count; i_it++ { 
            y = y_support.At(i_it)
            if (ToType(x.Id()).Contains(y) == CTRUE) { 
              y_out.AddFast(y)/*t=any,s=void*/
              } 
            } 
          } 
        Result = EID{y_out.Id(),0}
        } 
      }
      } 
    return Result} 
  
// The EID go function for: set! @ Param (throw: true) 
func E_set_I_Param (x EID) EID { 
    return F_set_I_Param(To_Param(OBJ(x)) )} 
  
/* {1} The go function for: size(x:Param) [status=1] */
func F_size_Param (x *ClaireParam ) EID { 
    var Result EID 
    { var arg_1 *ClaireSet  
      _ = arg_1
      var try_2 EID 
      /*g_try(v2:"try_2",loop:false) */
      try_2 = F_set_I_Param(x)
      /* ERROR PROTECTION INSERTED (arg_1-Result) */
      if ErrorIn(try_2) {Result = try_2
      } else {
      arg_1 = ToSet(OBJ(try_2))
      Result = EID{C__INT,IVAL(arg_1.Size())}
      }
      } 
    return Result} 
  
// The EID go function for: size @ Param (throw: true) 
func E_size_Param (x EID) EID { 
    return F_size_Param(To_Param(OBJ(x)) )} 
  
// subtype
/* {1} The go function for: set!(x:subtype) [status=1] */
func F_set_I_subtype (x *ClaireSubtype ) EID { 
    var Result EID 
    if (x.Arg.Id() == C_set.Id()) { 
      { var arg_1 *ClaireList  
        _ = arg_1
        var try_2 EID 
        /*g_try(v2:"try_2",loop:false) */
        { var arg_3 *ClaireAny  
          _ = arg_3
          var try_4 EID 
          /*g_try(v2:"try_4",loop:false) */
          try_4 = F_CALL(C_set_I,ARGS(EID{x.T1.Id(),0}))
          /* ERROR PROTECTION INSERTED (arg_3-try_2) */
          if ErrorIn(try_4) {try_2 = try_4
          } else {
          arg_3 = ANY(try_4)
          try_2 = EID{ToSet(arg_3).List_I().Id(),0}
          }
          } 
        /* ERROR PROTECTION INSERTED (arg_1-Result) */
        if ErrorIn(try_2) {Result = try_2
        } else {
        arg_1 = ToList(OBJ(try_2))
        Result = EID{F_build_powerset_list(arg_1).Id(),0}
        }
        } 
      } else {
      Result = ToException(C_general_error.Make(MakeString("[178] cannot enumerate ~S").Id(),MakeConstantList(x.Id()).Id())).Close()
      } 
    return Result} 
  
// The EID go function for: set! @ subtype (throw: true) 
func E_set_I_subtype (x EID) EID { 
    return F_set_I_subtype(ToSubtype(OBJ(x)) )} 
  
/* {1} The go function for: size(x:subtype) [status=1] */
func F_size_subtype (x *ClaireSubtype ) EID { 
    var Result EID 
    if (x.Arg.Id() == C_set.Id()) { 
      { var arg_1 *ClaireAny  
        _ = arg_1
        var try_2 EID 
        /*g_try(v2:"try_2",loop:false) */
        try_2 = F_CALL(C_size,ARGS(EID{x.T1.Id(),0}))
        /* ERROR PROTECTION INSERTED (arg_1-Result) */
        if ErrorIn(try_2) {Result = try_2
        } else {
        arg_1 = ANY(try_2)
        Result = F__exp2_integer(ToInteger(arg_1).Value)
        }
        } 
      } else {
      Result = ToException(C_general_error.Make(MakeString("[178] cannot enumerate ~S").Id(),MakeConstantList(x.Id()).Id())).Close()
      } 
    return Result} 
  
// The EID go function for: size @ subtype (throw: true) 
func E_size_subtype (x EID) EID { 
    return F_size_subtype(ToSubtype(OBJ(x)) )} 
  
// tuple
/* {1} The go function for: set!(x:tuple) [status=1] */
func F_set_I_tuple (x *ClaireTuple ) EID { 
    var Result EID 
    { var l *ClaireList   = ToList(x.Id())
      if (F_boolean_I_any(l.Id()).Id() != CTRUE.Id()) { 
        Result = EID{MakeConstantSet(CEMPTY.Id()).Id(),0}
        } else {
        { var l1 *ClaireSet  
          var try_1 EID 
          /*g_try(v2:"try_1",loop:false) */
          { var y_bag *ClaireSet   = ToType(CEMPTY.Id()).EmptySet()
            /*g_try(v2:"try_1",loop:false) */
            { 
              var y *ClaireAny  
              _ = y
              try_1= EID{CFALSE.Id(),0}
              var y_support *ClaireSet  
              var try_2 EID 
              /*g_try(v2:"try_2",loop:false) */
              try_2 = F_CALL(C_set_I,ARGS(l.At(1-1).ToEID()))
              /* ERROR PROTECTION INSERTED (y_support-try_1) */
              if ErrorIn(try_2) {try_1 = try_2
              } else {
              y_support = ToSet(OBJ(try_2))
              for i_it := 0; i_it < y_support.Count; i_it++ { 
                y = y_support.At(i_it)
                y_bag.AddFast(MakeConstantList(y).Id())/*t=any,s=void*/
                }
                } 
              } 
            /* ERROR PROTECTION INSERTED (try_1-try_1) */
            if !ErrorIn(try_1) {
            try_1 = EID{y_bag.Id(),0}
            }
            } 
          /* ERROR PROTECTION INSERTED (l1-Result) */
          if ErrorIn(try_1) {Result = try_1
          } else {
          l1 = ToSet(OBJ(try_1))
          /*g_try(v2:"Result",loop:true) */
          { var n int  = 2
            { var g0155 int  = l.Length()
              _ = g0155
              Result= EID{CFALSE.Id(),0}
              for (n <= g0155) { 
                /* While stat, v:"Result" loop:true */
                var loop_3 EID 
                _ = loop_3
                { 
                /*g_try(v2:"loop_3",loop:tuple("Result", EID)) */
                { var l2 *ClaireSet   = ToType(C_any.Id()).EmptySet()
                  _ = l2
                  /*g_try(v2:"loop_3",loop:tuple("Result", EID)) */
                  { 
                    var z *ClaireAny  
                    _ = z
                    loop_3= EID{CFALSE.Id(),0}
                    var z_support *ClaireSet  
                    var try_4 EID 
                    /*g_try(v2:"try_4",loop:false) */
                    try_4 = F_CALL(C_set_I,ARGS(l.At(n-1).ToEID()))
                    /* ERROR PROTECTION INSERTED (z_support-loop_3) */
                    if ErrorIn(try_4) {loop_3 = try_4
                    } else {
                    z_support = ToSet(OBJ(try_4))
                    for i_it := 0; i_it < z_support.Count; i_it++ { 
                      z = z_support.At(i_it)
                      { 
                        var l3 *ClaireList  
                        _ = l3
                        var l3_iter *ClaireAny  
                        var l3_support *ClaireSet  
                        l3_support = l1
                        for i_it := 0; i_it < l3_support.Count; i_it++ { 
                          l3_iter = l3_support.At(i_it)
                          l3 = ToList(l3_iter)
                          l2 = l2.AddFast(l3.Copy().AddFast(z).Id()/*t=any,s=any*/)/*t=any,s=set*/
                          } 
                        } 
                      }
                      } 
                    } 
                  /* ERROR PROTECTION INSERTED (loop_3-loop_3) */
                  if ErrorIn(loop_3) {Result = loop_3
                  break
                  } else {
                  l1 = l2
                  loop_3 = EID{l1.Id(),0}
                  }
                  } 
                /* ERROR PROTECTION INSERTED (loop_3-loop_3) */
                if ErrorIn(loop_3) {Result = loop_3
                break
                } else {
                n = (n+1)
                }
                /* try?:false, v2:"v_while7" loop will be:tuple("Result", EID) */
                } 
              }
              } 
            } 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Result = EID{l1.Id(),0}
          }
          }
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: set! @ tuple (throw: true) 
func E_set_I_tuple (x EID) EID { 
    return F_set_I_tuple(ToTuple(OBJ(x)) )} 
  
/* {1} The go function for: size(l:tuple) [status=1] */
func F_size_tuple (l *ClaireTuple ) EID { 
    var Result EID 
    if (F_boolean_I_any(l.Id()).Id() != CTRUE.Id()) { 
      Result = EID{C__INT,IVAL(1)}
      } else {
      { var m int 
        _ = m
        var try_1 EID 
        /*g_try(v2:"try_1",loop:false) */
        try_1 = F_CALL(C_size,ARGS(ToList(l.Id()).At(1-1).ToEID()))
        /* ERROR PROTECTION INSERTED (m-Result) */
        if ErrorIn(try_1) {Result = try_1
        } else {
        m = INT(try_1)
        /*g_try(v2:"Result",loop:true) */
        { var n int  = 2
          _ = n
          { var g0156 int  = l.Length()
            _ = g0156
            Result= EID{CFALSE.Id(),0}
            for (n <= g0156) { 
              /* While stat, v:"Result" loop:true */
              var loop_2 EID 
              _ = loop_2
              { 
              var try_3 EID 
              /*g_try(v2:"try_3",loop:tuple("Result", EID)) */
              { var arg_4 *ClaireAny  
                _ = arg_4
                var try_5 EID 
                /*g_try(v2:"try_5",loop:false) */
                try_5 = F_CALL(C_size,ARGS(ToList(l.Id()).At(n-1).ToEID()))
                /* ERROR PROTECTION INSERTED (arg_4-try_3) */
                if ErrorIn(try_5) {try_3 = try_5
                } else {
                arg_4 = ANY(try_5)
                try_3 = EID{C__INT,IVAL((m*ToInteger(arg_4).Value))}
                }
                } 
              /* ERROR PROTECTION INSERTED (m-loop_2) */
              if ErrorIn(try_3) {loop_2 = try_3
              Result = try_3
              break
              } else {
              m = INT(try_3)
              loop_2 = EID{C__INT,IVAL(m)}
              n = (n+1)
              }
              /* try?:false, v2:"v_while6" loop will be:tuple("Result", EID) */
              } 
            }
            } 
          } 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        Result = EID{C__INT,IVAL(m)}
        }
        }
        } 
      } 
    return Result} 
  
// The EID go function for: size @ tuple (throw: true) 
func E_size_tuple (l EID) EID { 
    return F_size_tuple(ToTuple(OBJ(l)) )} 
  
// declarations (now useless in CLAIRE4)
// ********************************************************************
// *                Part 3: Interface Methods                         *
// ********************************************************************
// there is a special restriction for + to specify the way the inheritance
// conflict should be solved
//U(self:set,ens:type) : type -> (case ens (set self /+ ens, any ens U self))
// the union makes a partial reduction to the normal form. The complete
// reduction is done by enumeration if needed during the type subsumption
// union is left-associative: A U B U C is represented by (A U B) U C  => never(t2(x:Union) % union)
// a union of intervals is ALWAYS disjoint
/* {1} The go function for: U(x:type,y:type) [status=0] */
func F_U_type (x *ClaireType ,y *ClaireType ) *ClaireType  { 
    // procedure body with s = type 
var Result *ClaireType  
    if (C_set.Id() == x.Isa.Id()) { 
      { var g0157 *ClaireSet   = ToSet(x.Id())
        if (C_set.Id() == y.Isa.Id()) { 
          { var g0158 *ClaireSet   = ToSet(y.Id())
            _ = g0158
            Result = ToType(F_append_set(g0157,g0158).Id())
            } 
          } else {
          Result = F_U_type(y,ToType(g0157.Id()))
          } 
        } 
      }  else if (y.Included(x) == CTRUE) { 
      Result = x
      }  else if (x.Included(y) == CTRUE) { 
      Result = y
      }  else if (y.Isa.IsIn(C_Union) == CTRUE) { 
      Result = F_U_type(F_U_type(x,ToType(OBJ(F_CALL(C_mClaire_t1,ARGS(EID{y.Id(),0}))))),To_Union(y.Id()).T2)
      } else {
      var g0166I *ClaireBoolean  
      if (x.Isa.IsIn(C_Interval) == CTRUE) { 
        g0166I = y.Isa.IsIn(C_Interval)
        } else {
        g0166I = CFALSE
        } 
      if (g0166I == CTRUE) { 
        if (((To_Interval(y.Id()).Arg1-1) <= To_Interval(x.Id()).Arg2) && 
            (To_Interval(x.Id()).Arg1 <= To_Interval(y.Id()).Arg1)) { 
          Result = F__dot_dot_integer(To_Interval(x.Id()).Arg1,To_Interval(y.Id()).Arg2)
          }  else if (((To_Interval(x.Id()).Arg1-1) <= To_Interval(y.Id()).Arg2) && 
            (To_Interval(y.Id()).Arg1 <= To_Interval(x.Id()).Arg1)) { 
          Result = F__dot_dot_integer(To_Interval(y.Id()).Arg1,To_Interval(x.Id()).Arg2)
          } else {
          { var _CL_obj *ClaireUnion   = To_Union(new(ClaireUnion).Is(C_Union))
            _CL_obj.T1 = x
            /*type->type*/_CL_obj.T2 = y
            /*type->type*/Result = ToType(_CL_obj.Id())
            } 
          } 
        } else {
        var g0167I *ClaireBoolean  
        if (x.Isa.IsIn(C_Union) == CTRUE) { 
          g0167I = y.Isa.IsIn(C_Interval)
          } else {
          g0167I = CFALSE
          } 
        if (g0167I == CTRUE) { 
          { var z *ClaireType   = F_U_type(To_Union(x.Id()).T2,y)
            if (z.Isa.IsIn(C_Union) == CTRUE) { 
              { var _CL_obj *ClaireUnion   = To_Union(new(ClaireUnion).Is(C_Union))
                _CL_obj.T1 = F_U_type(ToType(OBJ(F_CALL(C_mClaire_t1,ARGS(EID{x.Id(),0})))),y)
                /*type->type*/_CL_obj.T2 = To_Union(x.Id()).T2
                /*type->type*/Result = ToType(_CL_obj.Id())
                } 
              } else {
              Result = F_U_type(ToType(OBJ(F_CALL(C_mClaire_t1,ARGS(EID{x.Id(),0})))),z)
              } 
            } 
          } else {
          var g0168I *ClaireBoolean  
          if (x.Isa.IsIn(C_Interval) == CTRUE) { 
            { var g0165 *ClaireInterval   = To_Interval(x.Id())
              g0168I = MakeBoolean((C_set.Id() == y.Isa.Id()) && ((y.Contains(MakeInteger((g0165.Arg1-1)).Id()) == CTRUE) || 
                  (y.Contains(MakeInteger((g0165.Arg2+1)).Id()) == CTRUE)))
              } 
            } else {
            g0168I = CFALSE
            } 
          if (g0168I == CTRUE) { 
            { var a int  = To_Interval(x.Id()).Arg1
              { var b int  = To_Interval(x.Id()).Arg2
                if (y.Contains(MakeInteger((a-1)).Id()) == CTRUE) { 
                  a = (a-1)
                  } 
                if (y.Contains(MakeInteger((b+1)).Id()) == CTRUE) { 
                  b = (b+1)
                  } 
                Result = F_U_type(F__dot_dot_integer(a,b),y)
                } 
              } 
            } else {
            if (C_set.Id() == y.Isa.Id()) { 
              { var z_in *ClaireSet   = ToSet(y.Id())
                { var z_out *ClaireSet   = z_in.Empty()
                  { 
                    var z *ClaireAny  
                    _ = z
                    var z_support *ClaireSet  
                    z_support = z_in
                    for i_it := 0; i_it < z_support.Count; i_it++ { 
                      z = z_support.At(i_it)
                      if (x.Contains(z) != CTRUE) { 
                        z_out.AddFast(z)/*t=any,s=void*/
                        } 
                      } 
                    } 
                  y = ToType(z_out.Id())
                  } 
                } 
              } 
            { var _CL_obj *ClaireUnion   = To_Union(new(ClaireUnion).Is(C_Union))
              _CL_obj.T1 = x
              /*type->type*/_CL_obj.T2 = y
              /*type->type*/Result = ToType(_CL_obj.Id())
              } 
            } 
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: U @ type (throw: false) 
func E_U_type (x EID,y EID) EID { 
    return EID{F_U_type(ToType(OBJ(x)),ToType(OBJ(y)) ).Id(),0}} 
  
// the Interval construction method has a smart second-order type  - fix on v3.1.06
/* {1} The go function for: ..(x:integer,y:integer) [status=0] */
func F__dot_dot_integer (x int,y int) *ClaireType  { 
    // procedure body with s = type 
var Result *ClaireType  
    if (x <= y) { 
      Result = ToType(C_Interval.MakeInts(x,y))
      } else {
      Result = ToType(CEMPTY.Id())
      } 
    return Result} 
  
// The EID go function for: .. @ list<type_expression>(integer, integer) (throw: false) 
func E__dot_dot_integer (x EID,y EID) EID { 
    return EID{F__dot_dot_integer(INT(x),INT(y) ).Id(),0}} 
  
/* {1} The go function for: _dot_dot_integer_type */
func F__dot_dot_integer_type (x *ClaireType ,y *ClaireType ) EID { 
    var Result EID 
    var g0169I *ClaireBoolean  
    var try_1 EID 
    /*g_try(v2:"try_1",loop:false) */
    { 
      var v_and2 *ClaireBoolean  
      
      v_and2 = F_unique_ask_type(x)
      if (v_and2 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
      } else { 
        v_and2 = F_unique_ask_type(y)
        if (v_and2 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
        } else { 
          var try_2 EID 
          /*g_try(v2:"try_2",loop:false) */
          { var arg_3 *ClaireAny  
            _ = arg_3
            var try_5 EID 
            /*g_try(v2:"try_5",loop:false) */
            try_5 = F_the_type(x)
            /* ERROR PROTECTION INSERTED (arg_3-try_2) */
            if ErrorIn(try_5) {try_2 = try_5
            } else {
            arg_3 = ANY(try_5)
            { var arg_4 *ClaireAny  
              _ = arg_4
              var try_6 EID 
              /*g_try(v2:"try_6",loop:false) */
              try_6 = F_the_type(y)
              /* ERROR PROTECTION INSERTED (arg_4-try_2) */
              if ErrorIn(try_6) {try_2 = try_6
              } else {
              arg_4 = ANY(try_6)
              try_2 = F_CALL(ToProperty(C__inf_equal.Id()),ARGS(arg_3.ToEID(),arg_4.ToEID()))
              }
              } 
            }
            } 
          /* ERROR PROTECTION INSERTED (v_and2-try_1) */
          if ErrorIn(try_2) {try_1 = try_2
          } else {
          v_and2 = ToBoolean(OBJ(try_2))
          if (v_and2 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
          } else { 
            try_1 = EID{CTRUE.Id(),0}} 
          } 
        } 
      }
      } 
    /* ERROR PROTECTION INSERTED (g0169I-Result) */
    if ErrorIn(try_1) {Result = try_1
    } else {
    g0169I = ToBoolean(OBJ(try_1))
    if (g0169I == CTRUE) { 
      { 
        var v_bag_arg *ClaireAny  
        Result= EID{ToType(CEMPTY.Id()).EmptySet().Id(),0}
        var try_7 EID 
        /*g_try(v2:"try_7",loop:false) */
        { var arg_8 *ClaireAny  
          _ = arg_8
          var try_10 EID 
          /*g_try(v2:"try_10",loop:false) */
          try_10 = F_the_type(x)
          /* ERROR PROTECTION INSERTED (arg_8-try_7) */
          if ErrorIn(try_10) {try_7 = try_10
          } else {
          arg_8 = ANY(try_10)
          { var arg_9 *ClaireAny  
            _ = arg_9
            var try_11 EID 
            /*g_try(v2:"try_11",loop:false) */
            try_11 = F_the_type(y)
            /* ERROR PROTECTION INSERTED (arg_9-try_7) */
            if ErrorIn(try_11) {try_7 = try_11
            } else {
            arg_9 = ANY(try_11)
            try_7 = EID{F__dot_dot_integer(ToInteger(arg_8).Value,ToInteger(arg_9).Value).Id(),0}
            }
            } 
          }
          } 
        /* ERROR PROTECTION INSERTED (v_bag_arg-Result) */
        if ErrorIn(try_7) {Result = try_7
        } else {
        v_bag_arg = ANY(try_7)
        ToSet(OBJ(Result)).AddFast(v_bag_arg)}
        } 
      } else {
      Result = EID{F_nth_class1(C_subtype,ToType(C_integer.Id())).Id(),0}
      } 
    }
    return Result} 
  
  
// The dual EID go function for: "_dot_dot_integer_type" 
func E__dot_dot_integer_type (x EID,y EID) EID { 
    return F__dot_dot_integer_type(ToType(OBJ(x)),ToType(OBJ(y)))} 
  
// exception
/* {1} The go function for: but(s:any,x:any) [status=1] */
func F_but_any (s *ClaireAny ,x *ClaireAny ) EID { 
    var Result EID 
    if (s.Isa.IsIn(C_list) == CTRUE) { 
      { var g0170 *ClaireList   = ToList(s)
        _ = g0170
        { var y_in *ClaireList   = g0170
          { var y_out *ClaireList   = y_in.Empty()
            { 
              var y *ClaireAny  
              _ = y
              var y_support *ClaireList  
              y_support = y_in
              y_len := y_support.Length()
              for i_it := 0; i_it < y_len; i_it++ { 
                y = y_support.At(i_it)
                if (Equal(y,x) != CTRUE) { 
                  y_out.AddFast(y)/*t=any,s=void*/
                  } 
                } 
              } 
            Result = EID{y_out.Id(),0}
            } 
          } 
        } 
      }  else if (C_set.Id() == s.Isa.Id()) { 
      { var g0171 *ClaireSet   = ToSet(s)
        _ = g0171
        Result = EID{g0171.Copy().Delete(x).Id(),0}
        } 
      } else {
      { var arg_1 *ClaireList  
        _ = arg_1
        var try_2 EID 
        /*g_try(v2:"try_2",loop:false) */
        try_2 = F_enumerate_any(s)
        /* ERROR PROTECTION INSERTED (arg_1-Result) */
        if ErrorIn(try_2) {Result = try_2
        } else {
        arg_1 = ToList(OBJ(try_2))
        Result = EID{arg_1.Delete(x).Id(),0}
        }
        } 
      } 
    return Result} 
  
// The EID go function for: but @ any (throw: true) 
func E_but_any (s EID,x EID) EID { 
    return F_but_any(ANY(s),ANY(x) )} 
  
/* {1} The go function for: but_any_type */
func F_but_any_type (s *ClaireType ,x *ClaireType ) EID { 
    var Result EID 
    if (x.Included(ToType(C_list.Id())) == CTRUE) { 
      Result = EID{F_nth_class1(C_list,F_member_type(s)).Id(),0}
      }  else if (x.Included(ToType(C_set.Id())) == CTRUE) { 
      Result = EID{F_nth_class1(C_set,F_member_type(s)).Id(),0}
      } else {
      Result = EID{C_any.Id(),0}
      } 
    return Result} 
  
  
// The dual EID go function for: "but_any_type" 
func E_but_any_type (s EID,x EID) EID { 
    return F_but_any_type(ToType(OBJ(s)),ToType(OBJ(x)))} 
  
// a set difference (extended to types, with implicit enumeration)
/* {1} The go function for: \(x:type,y:type) [status=1] */
func F__backslash_type (x *ClaireType ,y *ClaireType ) EID { 
    var Result EID 
    { var z_out *ClaireSet   = ToType(CEMPTY.Id()).EmptySet()
      /*g_try(v2:"Result",loop:true) */
      { 
        var z *ClaireAny  
        _ = z
        Result= EID{CFALSE.Id(),0}
        var z_support *ClaireList  
        var try_1 EID 
        /*g_try(v2:"try_1",loop:false) */
        try_1 = F_enumerate_any(x.Id())
        /* ERROR PROTECTION INSERTED (z_support-Result) */
        if ErrorIn(try_1) {Result = try_1
        } else {
        z_support = ToList(OBJ(try_1))
        z_len := z_support.Length()
        for i_it := 0; i_it < z_len; i_it++ { 
          z = z_support.At(i_it)
          if (y.Contains(z) != CTRUE) { 
            z_out.AddFast(z)/*t=any,s=void*/
            } 
          }
          } 
        } 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = EID{z_out.Id(),0}
      }
      } 
    return Result} 
  
// The EID go function for: \ @ type (throw: true) 
func E__backslash_type (x EID,y EID) EID { 
    return F__backslash_type(ToType(OBJ(x)),ToType(OBJ(y)) )} 
  
// ******************************************************************
// *    Part 4: Lattice methods                                     *
// ******************************************************************
// glb operation ---------------------------------------------------
// gbl is the extension of the lattice operator ^ for types to type_expressions
// new in v3.0.60: we reintroduce a glb method
/* {1} The go function for: glb(x:set,y:type) [status=0] */
func F_glb_set (x *ClaireSet ,y *ClaireType ) *ClaireSet  { 
    // procedure body with s = set 
var Result *ClaireSet  
    { var z_in *ClaireSet   = x
      { var z_out *ClaireSet   = z_in.Empty()
        { 
          var z *ClaireAny  
          _ = z
          var z_support *ClaireSet  
          z_support = z_in
          for i_it := 0; i_it < z_support.Count; i_it++ { 
            z = z_support.At(i_it)
            if (y.Contains(z) == CTRUE) { 
              z_out.AddFast(z)/*t=any,s=void*/
              } 
            } 
          } 
        Result = z_out
        } 
      } 
    return Result} 
  
// The EID go function for: glb @ set (throw: false) 
func E_glb_set (x EID,y EID) EID { 
    return EID{F_glb_set(ToSet(OBJ(x)),ToType(OBJ(y)) ).Id(),0}} 
  
/* {1} The go function for: glb(x:Union,y:type) [status=0] */
func F_glb_Union (x *ClaireUnion ,y *ClaireType ) *ClaireType  { 
    return  F_U_type(ToType(OBJ(F_CALL(ToProperty(C_glb.Id()),ARGS(EID{x.T1.Id(),0},EID{y.Id(),0})))),ToType(OBJ(F_CALL(ToProperty(C_glb.Id()),ARGS(EID{x.T2.Id(),0},EID{y.Id(),0})))))
    } 
  
// The EID go function for: glb @ Union (throw: false) 
func E_glb_Union (x EID,y EID) EID { 
    return EID{F_glb_Union(To_Union(OBJ(x)),ToType(OBJ(y)) ).Id(),0}} 
  
/* {1} The go function for: glb(x:Interval,y:type) [status=0] */
func F_glb_Interval (x *ClaireInterval ,y *ClaireType ) *ClaireType  { 
    // procedure body with s = type 
var Result *ClaireType  
    if (C_class.Id() == y.Isa.Id()) { 
      { var g0173 *ClaireClass   = ToClass(y.Id())
        _ = g0173
        Result = ToType(IfThenElse((ToType(C_integer.Id()).Included(ToType(g0173.Id())) == CTRUE),
          x.Id(),
          CEMPTY.Id()))
        } 
      }  else if (C_set.Id() == y.Isa.Id()) { 
      { var g0174 *ClaireSet   = ToSet(y.Id())
        _ = g0174
        Result = ToType(F_glb_set(g0174,ToType(x.Id())).Id())
        } 
      }  else if (y.Isa.IsIn(C_Interval) == CTRUE) { 
      { var g0175 *ClaireInterval   = To_Interval(y.Id())
        if (x.Arg1 <= g0175.Arg1) { 
          if (g0175.Arg1 <= x.Arg2) { 
            if (x.Arg2 <= g0175.Arg2) { 
              Result = F__dot_dot_integer(g0175.Arg1,x.Arg2)
              } else {
              Result = ToType(g0175.Id())
              } 
            } else {
            Result = ToType(CEMPTY.Id())
            } 
          } else {
          Result = F_glb_Interval(g0175,ToType(x.Id()))
          } 
        } 
      }  else if (y.Isa.IsIn(C_Union) == CTRUE) { 
      { var g0176 *ClaireUnion   = To_Union(y.Id())
        Result = F_U_type(F_glb_Interval(x,g0176.T1),F_glb_Interval(x,g0176.T2))
        } 
      } else {
      Result = ToType(CEMPTY.Id())
      } 
    return Result} 
  
// The EID go function for: glb @ Interval (throw: false) 
func E_glb_Interval (x EID,y EID) EID { 
    return EID{F_glb_Interval(To_Interval(OBJ(x)),ToType(OBJ(y)) ).Id(),0}} 
  
/* {1} The go function for: glb(x:class,y:type) [status=0] */
func F_glb_class (x *ClaireClass ,y *ClaireType ) *ClaireType  { 
    // procedure body with s = type 
var Result *ClaireType  
    if ((x.Open == ClEnv.ABSTRACT) && 
        (F_boolean_I_any(x.Subclass.Id()).Id() != CTRUE.Id())) { 
      { var z_out *ClaireSet   = ToType(CEMPTY.Id()).EmptySet()
        { 
          var z *ClaireAny  
          _ = z
          var z_support *ClaireList  
          z_support = x.Instances
          z_len := z_support.Length()
          for i_it := 0; i_it < z_len; i_it++ { 
            z = z_support.At(i_it)
            if (y.Contains(z) == CTRUE) { 
              z_out.AddFast(z)/*t=any,s=void*/
              } 
            } 
          } 
        Result = ToType(z_out.Id())
        } 
      }  else if ((x.Open == ClEnv.ABSTRACT) && 
        (F_boolean_I_any(x.Instances.Id()).Id() != CTRUE.Id())) { 
      { var arg_1 *ClaireList  
        _ = arg_1
        { 
          var v_list4 *ClaireSet  
          var z *ClaireType  
          var v_local4 *ClaireAny  
          v_list4 = x.Subclass
          arg_1 = CreateList(ToType(CEMPTY.Id()),v_list4.Length())
          for CLcount := 0; CLcount < v_list4.Count; CLcount++{ 
            z = ToType(v_list4.At(CLcount))
            v_local4 = ANY(F_CALL(ToProperty(C_glb.Id()),ARGS(EID{z.Id(),0},EID{y.Id(),0})))
            arg_1.PutAt(CLcount,v_local4)
            } 
          } 
        Result = F_Uall_list(arg_1)
        } 
      }  else if (C_class.Id() == y.Isa.Id()) { 
      { var g0178 *ClaireClass   = ToClass(y.Id())
        _ = g0178
        Result = F_join_class(x,g0178)
        } 
      } else {
      Result = ToType(OBJ(F_CALL(ToProperty(C_glb.Id()),ARGS(EID{y.Id(),0},EID{x.Id(),0}))))
      } 
    return Result} 
  
// The EID go function for: glb @ class (throw: false) 
func E_glb_class (x EID,y EID) EID { 
    return EID{F_glb_class(ToClass(OBJ(x)),ToType(OBJ(y)) ).Id(),0}} 
  
/* {1} The go function for: glb(x:Param,y:type) [status=0] */
func F_glb_Param (x *ClaireParam ,y *ClaireType ) *ClaireType  { 
    // procedure body with s = type 
var Result *ClaireType  
    if (y.Isa.IsIn(C_Param) == CTRUE) { 
      { var g0180 *ClaireParam   = To_Param(y.Id())
        { var c *ClaireType   = F_join_class(x.Arg,g0180.Arg)
          { var lp *ClaireList   = x.Params.Append(g0180.Params).Set_I().List_I()
            { var l *ClaireList   = ToType(C_any.Id()).EmptyList()
              _ = l
              { 
                var p *ClaireAny  
                _ = p
                var p_support *ClaireList  
                p_support = lp
                p_len := p_support.Length()
                for i_it := 0; i_it < p_len; i_it++ { 
                  p = p_support.At(i_it)
                  { var t *ClaireType   = ToType(OBJ(F_CALL(ToProperty(C_glb.Id()),ARGS(EID{x.At(ToProperty(p)).Id(),0},EID{g0180.At(ToProperty(p)).Id(),0}))))
                    _ = t
                    if (Equal(t.Id(),CEMPTY.Id()) != CTRUE) { 
                      l = l.AddFast(t.Id())/*t=any,s=list*/
                      } else {
                      c = ToType(CEMPTY.Id())
                      
                      break
                      } 
                    } 
                  } 
                } 
              if (Equal(c.Id(),CEMPTY.Id()) != CTRUE) { 
                { var _CL_obj *ClaireParam   = To_Param(new(ClaireParam).Is(C_Param))
                  _CL_obj.Arg = ToClass(c.Id())
                  /*class->class*/_CL_obj.Params = lp
                  /*list->list*/_CL_obj.Args = l
                  /*list->list*/Result = ToType(_CL_obj.Id())
                  } 
                } else {
                Result = ToType(CEMPTY.Id())
                } 
              } 
            } 
          } 
        } 
      }  else if (C_class.Id() == y.Isa.Id()) { 
      { var g0181 *ClaireClass   = ToClass(y.Id())
        _ = g0181
        { var c *ClaireType   = F_join_class(x.Arg,g0181)
          if (Equal(c.Id(),CEMPTY.Id()) != CTRUE) { 
            { var _CL_obj *ClaireParam   = To_Param(new(ClaireParam).Is(C_Param))
              _CL_obj.Arg = ToClass(c.Id())
              /*class->class*/_CL_obj.Params = x.Params
              /*list->list*/_CL_obj.Args = x.Args
              /*list->list*/Result = ToType(_CL_obj.Id())
              } 
            } else {
            Result = ToType(CEMPTY.Id())
            } 
          } 
        } 
      } else {
      Result = ToType(OBJ(F_CALL(ToProperty(C_glb.Id()),ARGS(EID{y.Id(),0},EID{x.Id(),0}))))
      } 
    return Result} 
  
// The EID go function for: glb @ Param (throw: false) 
func E_glb_Param (x EID,y EID) EID { 
    return EID{F_glb_Param(To_Param(OBJ(x)),ToType(OBJ(y)) ).Id(),0}} 
  
// notice that a param whose class is a type must use of (only parameter allowed!)
// the result is a subtype
/* {1} The go function for: glb(x:subtype,y:type) [status=0] */
func F_glb_subtype (x *ClaireSubtype ,y *ClaireType ) *ClaireType  { 
    // procedure body with s = type 
var Result *ClaireType  
    if (C_class.Id() == y.Isa.Id()) { 
      { var g0183 *ClaireClass   = ToClass(y.Id())
        if (Equal(F_join_class(x.Arg,g0183).Id(),CEMPTY.Id()) != CTRUE) { 
          Result = F_nth_class1(ToClass(F_join_class(x.Arg,g0183).Id()),x.T1)
          } else {
          Result = ToType(CEMPTY.Id())
          } 
        } 
      }  else if (y.Isa.IsIn(C_Param) == CTRUE) { 
      { var g0184 *ClaireParam   = To_Param(y.Id())
        if (Equal(F_join_class(x.Arg,g0184.Arg).Id(),CEMPTY.Id()) != CTRUE) { 
          Result = F_param_I_class(ToClass(F_join_class(x.Arg,g0184.Arg).Id()),ToType(OBJ(F_CALL(ToProperty(C_glb.Id()),ARGS(EID{F_member_type(ToType(x.Id())).Id(),0},EID{F_member_type(ToType(g0184.Id())).Id(),0})))))
          } else {
          Result = ToType(CEMPTY.Id())
          } 
        } 
      }  else if (y.Isa.IsIn(C_subtype) == CTRUE) { 
      { var g0185 *ClaireSubtype   = ToSubtype(y.Id())
        if (Equal(F_join_class(x.Arg,g0185.Arg).Id(),CEMPTY.Id()) != CTRUE) { 
          { var t *ClaireType   = ToType(OBJ(F_CALL(ToProperty(C_glb.Id()),ARGS(EID{x.T1.Id(),0},EID{g0185.T1.Id(),0}))))
            if (Equal(t.Id(),CEMPTY.Id()) != CTRUE) { 
              Result = F_nth_class1(ToClass(F_join_class(x.Arg,g0185.Arg).Id()),t)
              } else {
              Result = ToType(CEMPTY.Id())
              } 
            } 
          } else {
          Result = ToType(CEMPTY.Id())
          } 
        } 
      } else {
      Result = ToType(OBJ(F_CALL(ToProperty(C_glb.Id()),ARGS(EID{y.Id(),0},EID{x.Id(),0}))))
      } 
    return Result} 
  
// The EID go function for: glb @ subtype (throw: false) 
func E_glb_subtype (x EID,y EID) EID { 
    return EID{F_glb_subtype(ToSubtype(OBJ(x)),ToType(OBJ(y)) ).Id(),0}} 
  
// set, Interval, list
/* {1} The go function for: glb(x:tuple,y:type) [status=0] */
func F_glb_tuple (x *ClaireTuple ,y *ClaireType ) *ClaireType  { 
    // procedure body with s = type 
var Result *ClaireType  
    if (C_class.Id() == y.Isa.Id()) { 
      { var g0187 *ClaireClass   = ToClass(y.Id())
        _ = g0187
        Result = ToType(IfThenElse((C_tuple.IsIn(g0187) == CTRUE),
          x.Id(),
          CEMPTY.Id()))
        } 
      }  else if (y.Isa.IsIn(C_Param) == CTRUE) { 
      Result = ToType(CEMPTY.Id())
      }  else if (C_tuple.Id() == y.Isa.Id()) { 
      { var g0189 *ClaireTuple   = ToTuple(y.Id())
        _ = g0189
        Result = ToType(F__exp_list(ToList(x.Id()),ToList(g0189.Id())).Tuple_I().Id())
        } 
      }  else if (y.Isa.IsIn(C_subtype) == CTRUE) { 
      { var g0190 *ClaireSubtype   = ToSubtype(y.Id())
        if (g0190.Arg.Id() == C_tuple.Id()) { 
          { var arg_1 *ClaireList  
            _ = arg_1
            { var z_bag *ClaireList   = ToType(CEMPTY.Id()).EmptyList()
              { 
                var z *ClaireAny  
                _ = z
                var z_support *ClaireList  
                z_support = ToList(x.Id())
                z_len := z_support.Length()
                for i_it := 0; i_it < z_len; i_it++ { 
                  z = z_support.At(i_it)
                  z_bag.AddFast(ANY(F_CALL(ToProperty(C_glb.Id()),ARGS(z.ToEID(),EID{g0190.T1.Id(),0}))))/*t=any,s=void*/
                  } 
                } 
              arg_1 = z_bag
              } 
            Result = ToType(arg_1.Tuple_I().Id())
            } 
          } else {
          Result = ToType(CEMPTY.Id())
          } 
        } 
      } else {
      Result = ToType(OBJ(F_CALL(ToProperty(C_glb.Id()),ARGS(EID{y.Id(),0},EID{x.Id(),0}))))
      } 
    return Result} 
  
// The EID go function for: glb @ tuple (throw: false) 
func E_glb_tuple (x EID,y EID) EID { 
    return EID{F_glb_tuple(ToTuple(OBJ(x)),ToType(OBJ(y)) ).Id(),0}} 
  
// a reference is seen as "any"
/* {1} The go function for: glb(x:Reference,y:type) [status=0] */
func F_glb_Reference (x *ClaireReference ,y *ClaireType ) *ClaireType  { 
    return  y
    } 
  
// The EID go function for: glb @ Reference (throw: false) 
func E_glb_Reference (x EID,y EID) EID { 
    return EID{F_glb_Reference(To_Reference(OBJ(x)),ToType(OBJ(y)) ).Id(),0}} 
  
// this will be greatly simplified in a few minutes !
/* {1} The go function for: ^(x:type,y:type) [status=0] */
func F__exp_type (x *ClaireType ,y *ClaireType ) *ClaireType  { 
    return  ToType(OBJ(F_CALL(ToProperty(C_glb.Id()),ARGS(EID{x.Id(),0},EID{y.Id(),0}))))
    } 
  
// The EID go function for: ^ @ type (throw: false) 
func E__exp_type (x EID,y EID) EID { 
    return EID{F__exp_type(ToType(OBJ(x)),ToType(OBJ(y)) ).Id(),0}} 
  
// the old lattice_glb
/* {1} The go function for: join(x:class,y:class) [status=0] */
func F_join_class (x *ClaireClass ,y *ClaireClass ) *ClaireType  { 
    // procedure body with s = type 
var Result *ClaireType  
    { var l1 *ClaireList   = x.Ancestors
      { var n1 int  = l1.Length()
        { var l2 *ClaireList   = y.Ancestors
          { var n2 int  = l2.Length()
            if (n1 < n2) { 
              Result = ToType(IfThenElse((l2.ValuesO()[n1-1] == x.Id()),
                y.Id(),
                CEMPTY.Id()))
              }  else if (l1.ValuesO()[n2-1] == y.Id()) { 
              Result = ToType(x.Id())
              } else {
              Result = ToType(CEMPTY.Id())
              } 
            } 
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: join @ class (throw: false) 
func E_join_class (x EID,y EID) EID { 
    return EID{F_join_class(ToClass(OBJ(x)),ToClass(OBJ(y)) ).Id(),0}} 
  
// for lists
/* {1} The go function for: ^(x:list,y:list) [status=0] */
func F__exp_list (x *ClaireList ,y *ClaireList ) *ClaireList  { 
    // procedure body with s = list 
var Result *ClaireList  
    { var n int  = x.Length()
      { var r *ClaireList   = ToType(CEMPTY.Id()).EmptyList()
        _ = r
        if (n == y.Length()) { 
          { var i int  = 1
            { var g0192 int  = n
              _ = g0192
              for (i <= g0192) { 
                /* While stat, v:"Result" loop:false */
                { var z *ClaireTypeExpression   = ToTypeExpression(OBJ(F_CALL(ToProperty(C_glb.Id()),ARGS(x.At(i-1).ToEID(),y.At(i-1).ToEID()))))
                  _ = z
                  if (Equal(z.Id(),CEMPTY.Id()) != CTRUE) { 
                    r = r.AddFast(z.Id())/*t=any,s=list*/
                    } else {
                    r = CNIL
                    
                    break
                    } 
                  } 
                i = (i+1)
                /* try?:false, v2:"v_while7" loop will be:tuple("Result", void) */
                } 
              } 
            } 
          } 
        Result = r
        } 
      } 
    return Result} 
  
// The EID go function for: ^ @ list (throw: false) 
func E__exp_list (x EID,y EID) EID { 
    return EID{F__exp_list(ToList(OBJ(x)),ToList(OBJ(y)) ).Id(),0}} 
  
// a combined union
/* {1} The go function for: Uall(l:list) [status=0] */
func F_Uall_list (l *ClaireList ) *ClaireType  { 
    // procedure body with s = type 
var Result *ClaireType  
    { var rep *ClaireSet   = CEMPTY
      _ = rep
      { 
        var x *ClaireAny  
        _ = x
        var x_support *ClaireList  
        x_support = l
        x_len := x_support.Length()
        for i_it := 0; i_it < x_len; i_it++ { 
          x = x_support.At(i_it)
          rep = ToSet(F_U_type(ToType(rep.Id()),ToType(x)).Id())
          } 
        } 
      Result = ToType(rep.Id())
      } 
    return Result} 
  
// The EID go function for: Uall @ list (throw: false) 
func E_Uall_list (l EID) EID { 
    return EID{F_Uall_list(ToList(OBJ(l)) ).Id(),0}} 
  
// ------------------- The inclusion operation ------------------------
// the specialized versions %t and <=t are hard coded in Kernel, hence not extensible.
// if we create new types they will be used as patterns, not concrete types.
// hand-made
// v4 open coded (link to Included kernel method)
/* {1} The go function for: <=t(s:type,y:type) [status=0] */
func F__inf_equalt_type (s *ClaireType ,y *ClaireType ) *ClaireBoolean  { 
    if (s.Included(y) == CTRUE) {return CTRUE
    } else {return CFALSE}} 
  
// The EID go function for: <=t @ type (throw: false) 
func E__inf_equalt_type (s EID,y EID) EID { 
    return EID{F__inf_equalt_type(ToType(OBJ(s)),ToType(OBJ(y)) ).Id(),0}} 
  
// default order for types
/* {1} The go function for: <=(x:type_expression,y:type_expression) [status=1] */
func F__inf_equal_type_expression (x *ClaireTypeExpression ,y *ClaireTypeExpression ) EID { 
    var Result EID 
    if (C_set.Id() == x.Isa.Id()) { 
      { var g0193 *ClaireSet   = ToSet(x.Id())
        _ = g0193
        { var arg_1 *ClaireAny  
          _ = arg_1
          var try_2 EID 
          /*g_try(v2:"try_2",loop:false) */
          { 
            var z *ClaireAny  
            _ = z
            try_2= EID{CFALSE.Id(),0}
            var z_support *ClaireSet  
            z_support = g0193
            for i_it := 0; i_it < z_support.Count; i_it++ { 
              z = z_support.At(i_it)
              var loop_3 EID 
              _ = loop_3
              /*g_try(v2:"loop_3",loop:tuple("try_2", EID)) */
              var g0198I *ClaireBoolean  
              var try_4 EID 
              /*g_try(v2:"try_4",loop:false) */
              { var arg_5 *ClaireBoolean  
                _ = arg_5
                var try_6 EID 
                /*g_try(v2:"try_6",loop:false) */
                try_6 = F_BELONG(z,y.Id())
                /* ERROR PROTECTION INSERTED (arg_5-try_4) */
                if ErrorIn(try_6) {try_4 = try_6
                } else {
                arg_5 = ToBoolean(OBJ(try_6))
                try_4 = EID{arg_5.Not.Id(),0}
                }
                } 
              /* ERROR PROTECTION INSERTED (g0198I-loop_3) */
              if ErrorIn(try_4) {loop_3 = try_4
              } else {
              g0198I = ToBoolean(OBJ(try_4))
              if (g0198I == CTRUE) { 
                try_2 = EID{CTRUE.Id(),0}
                break
                } else {
                loop_3 = EID{CFALSE.Id(),0}
                } 
              }
              /* ERROR PROTECTION INSERTED (loop_3-try_2) */
              if ErrorIn(loop_3) {try_2 = loop_3
              break
              } else {
              }
              } 
            } 
          /* ERROR PROTECTION INSERTED (arg_1-Result) */
          if ErrorIn(try_2) {Result = try_2
          } else {
          arg_1 = ANY(try_2)
          Result = EID{F_not_any(arg_1).Id(),0}
          }
          } 
        } 
      }  else if (x.Isa.IsIn(C_type) == CTRUE) { 
      { var g0194 *ClaireType   = ToType(x.Id())
        if (y.Isa.IsIn(C_type) == CTRUE) { 
          { var g0195 *ClaireType   = ToType(y.Id())
            _ = g0195
            Result = EID{g0194.Included(g0195).Id(),0}
            } 
          } else {
          { var z *ClaireAny   = g0194.Id()
            _ = z
            Result = F_CALL(ToProperty(C_less_ask.Id()),ARGS(z.ToEID(),EID{y.Id(),0}))
            } 
          } 
        } 
      } else {
      Result = F_CALL(ToProperty(C_less_ask.Id()),ARGS(EID{x.Id(),0},EID{y.Id(),0}))
      } 
    return Result} 
  
// The EID go function for: <= @ type_expression (throw: true) 
func E__inf_equal_type_expression (x EID,y EID) EID { 
    return F__inf_equal_type_expression(ToTypeExpression(OBJ(x)),ToTypeExpression(OBJ(y)) )} 
  
// membership for types
// hand-made
// v4 open coded (link to Contains kernel method)
/* {1} The go function for: %t(x:any,y:type) [status=0] */
func F_Core__Zt_any (x *ClaireAny ,y *ClaireType ) *ClaireBoolean  { 
    if (y.Contains(x) == CTRUE) {return CTRUE
    } else {return CFALSE}} 
  
// The EID go function for: %t @ any (throw: false) 
func E_Core__Zt_any (x EID,y EID) EID { 
    return EID{F_Core__Zt_any(ANY(x),ToType(OBJ(y)) ).Id(),0}} 
  
// extensibility for type_expression is through less?, that always returns a value (hence no error returned)
/* {1} The go function for: less?(x:type_expression,y:type_expression) [status=0] */
func F_less_ask_type_expression (x *ClaireTypeExpression ,y *ClaireTypeExpression ) *ClaireBoolean  { 
    if (CFALSE == CTRUE) {return CTRUE
    } else {return CFALSE}} 
  
// The EID go function for: less? @ list<type_expression>(type_expression, type_expression) (throw: false) 
func E_less_ask_type_expression (x EID,y EID) EID { 
    return EID{F_less_ask_type_expression(ToTypeExpression(OBJ(x)),ToTypeExpression(OBJ(y)) ).Id(),0}} 
  
// ******************************************************************
// *    Part 5: type methods                                        *
// ******************************************************************
// --------------------- extract tuple type information -------------
// extract a member type, that is a valid type for all members (z) of instances of
// the type x.This is much simpler in v3.0
/* {1} The go function for: member(x:type) [status=0] */
func F_member_type (x *ClaireType ) *ClaireType  { 
    // procedure body with s = type 
var Result *ClaireType  
    if (C_class.Id() == x.Isa.Id()) { 
      { var g0199 *ClaireClass   = ToClass(x.Id())
        _ = g0199
        if (g0199.Id() == C_Interval.Id()) { 
          Result = ToType(C_integer.Id())
          } else {
          Result = ToType(C_any.Id())
          } 
        } 
      }  else if (x.Isa.IsIn(C_Union) == CTRUE) { 
      { var g0200 *ClaireUnion   = To_Union(x.Id())
        Result = F_U_type(F_member_type(g0200.T1),F_member_type(g0200.T2))
        } 
      }  else if (x.Isa.IsIn(C_Interval) == CTRUE) { 
      Result = ToType(CEMPTY.Id())
      }  else if (x.Isa.IsIn(C_Param) == CTRUE) { 
      { var g0202 *ClaireParam   = To_Param(x.Id())
        _ = g0202
        Result = F_member_type(g0202.At(C_of))
        } 
      }  else if (C_tuple.Id() == x.Isa.Id()) { 
      { var g0203 *ClaireTuple   = ToTuple(x.Id())
        _ = g0203
        Result = F_Uall_list(ToList(g0203.Id()))
        } 
      }  else if (x.Isa.IsIn(C_subtype) == CTRUE) { 
      { var g0204 *ClaireSubtype   = ToSubtype(x.Id())
        _ = g0204
        Result = g0204.T1
        } 
      }  else if (C_set.Id() == x.Isa.Id()) { 
      { var g0205 *ClaireSet   = ToSet(x.Id())
        _ = g0205
        { var arg_1 *ClaireList  
          _ = arg_1
          { 
            var v_list5 *ClaireSet  
            var y *ClaireAny  
            var v_local5 *ClaireAny  
            v_list5 = g0205
            arg_1 = CreateList(ToType(CEMPTY.Id()),v_list5.Length())
            for CLcount := 0; CLcount < v_list5.Count; CLcount++{ 
              y = v_list5.At(CLcount)
              if (y.Isa.IsIn(C_list) == CTRUE) { 
                { var g0206 *ClaireList   = ToList(y)
                  _ = g0206
                  v_local5 = g0206.Set_I().Id()
                  } 
                }  else if (y.Isa.IsIn(C_type) == CTRUE) { 
                { var g0207 *ClaireType   = ToType(y)
                  _ = g0207
                  v_local5 = g0207.Id()
                  } 
                } else {
                v_local5 = CEMPTY.Id()
                } 
              arg_1.PutAt(CLcount,v_local5)
              } 
            } 
          Result = F_Uall_list(arg_1)
          } 
        } 
      } else {
      Result = ToType(CEMPTY.Id())
      } 
    return Result} 
  
// The EID go function for: member @ type (throw: false) 
func E_member_type (x EID) EID { 
    return EID{F_member_type(ToType(OBJ(x)) ).Id(),0}} 
  
// a simpler version (projection on bag subtypes)
// dumb code because it is used early in the bootstrap
/* {1} The go function for: of_extract(x:type) [status=0] */
func F_of_extract_type (x *ClaireType ) *ClaireType  { 
    // procedure body with s = type 
var Result *ClaireType  
    { var c *ClaireClass   = x.Isa
      if (c.Id() == C_subtype.Id()) { 
        Result = ToSubtype(x.Id()).T1
        }  else if (c.Id() == C_Param.Id()) { 
        if (To_Param(x.Id()).Params.At(1-1) == C_of.Id()) { 
          { var y *ClaireType   = ToType(To_Param(x.Id()).Args.At(1-1))
            if (C_set.Id() == y.Isa.Id()) { 
              { var g0210 *ClaireSet   = ToSet(y.Id())
                _ = g0210
                Result = ToType(g0210.List_I().At(1-1))
                } 
              }  else if (y.Isa.IsIn(C_subtype) == CTRUE) { 
              { var g0211 *ClaireSubtype   = ToSubtype(y.Id())
                _ = g0211
                Result = g0211.T1
                } 
              } else {
              Result = ToType(C_any.Id())
              } 
            } 
          } else {
          Result = ToType(C_any.Id())
          } 
        } else {
        Result = ToType(C_any.Id())
        } 
      } 
    return Result} 
  
// The EID go function for: of_extract @ type (throw: false) 
func E_of_extract_type (x EID) EID { 
    return EID{F_of_extract_type(ToType(OBJ(x)) ).Id(),0}} 
  
// useful type functions for the compiler
/* {1} The go function for: unique?(x:type) [status=0] */
func F_unique_ask_type (x *ClaireType ) *ClaireBoolean  { 
    // procedure body with s = boolean 
var Result *ClaireBoolean  
    if (C_set.Id() == x.Isa.Id()) { 
      { var g0213 *ClaireSet   = ToSet(x.Id())
        _ = g0213
        Result = Equal(MakeInteger(g0213.Size()).Id(),MakeInteger(1).Id())
        } 
      }  else if (C_class.Id() == x.Isa.Id()) { 
      { var g0214 *ClaireClass   = ToClass(x.Id())
        Result = MakeBoolean((g0214.Open == 0) && (F_size_class(g0214) == 1))
        } 
      } else {
      Result = CFALSE
      } 
    return Result} 
  
// The EID go function for: unique? @ type (throw: false) 
func E_unique_ask_type (x EID) EID { 
    return EID{F_unique_ask_type(ToType(OBJ(x)) ).Id(),0}} 
  
// returns the unique element of the type
/* {1} The go function for: the(x:type) [status=1] */
func F_the_type (x *ClaireType ) EID { 
    var Result EID 
    { var arg_1 *ClaireList  
      _ = arg_1
      var try_2 EID 
      /*g_try(v2:"try_2",loop:false) */
      { var arg_3 *ClaireAny  
        _ = arg_3
        var try_4 EID 
        /*g_try(v2:"try_4",loop:false) */
        try_4 = F_CALL(C_set_I,ARGS(EID{x.Id(),0}))
        /* ERROR PROTECTION INSERTED (arg_3-try_2) */
        if ErrorIn(try_4) {try_2 = try_4
        } else {
        arg_3 = ANY(try_4)
        try_2 = EID{ToSet(arg_3).List_I().Id(),0}
        }
        } 
      /* ERROR PROTECTION INSERTED (arg_1-Result) */
      if ErrorIn(try_2) {Result = try_2
      } else {
      arg_1 = ToList(OBJ(try_2))
      Result = arg_1.At(1-1).ToEID()
      }
      } 
    return Result} 
  
// The EID go function for: the @ type (throw: true) 
func E_the_type (x EID) EID { 
    return F_the_type(ToType(OBJ(x)) )} 
  
// bitvector made easy
// v0.01: should not use set[0 .. 29] => burden on caller is too heavy
/* {1} The go function for: integer!(s:set[integer]) [status=1] */
func F_integer_I_set (s *ClaireSet ) EID { 
    var Result EID 
    { var n int  = 0
      _ = n
      /*g_try(v2:"Result",loop:true) */
      { 
        var y int 
        _ = y
        var y_iter *ClaireAny  
        Result= EID{CFALSE.Id(),0}
        var y_support *ClaireSet  
        y_support = s
        for i_it := 0; i_it < y_support.Count; i_it++ { 
          y_iter = y_support.At(i_it)
          y = ToInteger(y_iter).Value
          var loop_1 EID 
          _ = loop_1
          /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
          if ((y >= 0) && 
              (y <= 29)) { 
            var try_2 EID 
            /*g_try(v2:"try_2",loop:false) */
            { var arg_3 int 
              _ = arg_3
              var try_4 EID 
              /*g_try(v2:"try_4",loop:false) */
              try_4 = F__exp2_integer(y)
              /* ERROR PROTECTION INSERTED (arg_3-try_2) */
              if ErrorIn(try_4) {try_2 = try_4
              } else {
              arg_3 = INT(try_4)
              try_2 = EID{C__INT,IVAL((n+arg_3))}
              }
              } 
            /* ERROR PROTECTION INSERTED (n-loop_1) */
            if ErrorIn(try_2) {loop_1 = try_2
            } else {
            n = INT(try_2)
            loop_1 = EID{C__INT,IVAL(n)}
            }
            } else {
            loop_1 = EID{CFALSE.Id(),0}
            } 
          /* ERROR PROTECTION INSERTED (loop_1-Result) */
          if ErrorIn(loop_1) {Result = loop_1
          break
          } else {
          }
          } 
        } 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = EID{C__INT,IVAL(n)}
      }
      } 
    return Result} 
  
// The EID go function for: integer! @ set (throw: true) 
func E_integer_I_set (s EID) EID { 
    return F_integer_I_set(ToSet(OBJ(s)) )} 
  
/* {1} The go function for: make_set(x:integer) [status=0] */
func F_make_set_integer (x int) *ClaireSet  { 
    // procedure body with s = set 
var Result *ClaireSet  
    { var i_out *ClaireSet   = ToType(CEMPTY.Id()).EmptySet()
      { var i int  = 0
        { var g0216 int  = 29
          _ = g0216
          for (i <= g0216) { 
            /* While stat, v:"Result" loop:false */
            if (F_nth_integer(x,i) == CTRUE) { 
              i_out.AddFast(MakeInteger(i).Id())/*t=any,s=void*/
              } 
            i = (i+1)
            /* try?:false, v2:"v_while5" loop will be:tuple("Result", void) */
            } 
          } 
        } 
      Result = i_out
      } 
    return Result} 
  
// The EID go function for: make_set @ integer (throw: false) 
func E_make_set_integer (x EID) EID { 
    return EID{F_make_set_integer(INT(x) ).Id(),0}} 
  
// asbtract coercion of a set into an interval
/* {1} The go function for: abstract_type(xt1:set) [status=0] */
func F_abstract_type_set (xt1 *ClaireSet ) *ClaireType  { 
    // procedure body with s = type 
var Result *ClaireType  
    { var m1 int  = 1
      { var m2 int  = 0
        { 
          var x *ClaireAny  
          _ = x
          var x_support *ClaireSet  
          x_support = xt1
          for i_it := 0; i_it < x_support.Count; i_it++ { 
            x = x_support.At(i_it)
            if (C_integer.Id() == x.Isa.Id()) { 
              { var g0217 int  = ToInteger(x).Value
                if (m1 > m2) { 
                  m1 = g0217
                  m2 = g0217
                  }  else if (g0217 > m2) { 
                  m2 = g0217
                  }  else if (g0217 < m1) { 
                  m1 = g0217
                  } 
                } 
              } else {
              m1 = 1
              m2 = 0
              
              break
              } 
            } 
          } 
        Result = F__dot_dot_integer(m1,m2)
        } 
      } 
    return Result} 
  
// The EID go function for: abstract_type @ set (throw: false) 
func E_abstract_type_set (xt1 EID) EID { 
    return EID{F_abstract_type_set(ToSet(OBJ(xt1)) ).Id(),0}} 
  
// abstract interpretation of integer arithmetique
/* {1} The go function for: abstract_type(p:operation,xt1:type,xt2:type) [status=0] */
func F_abstract_type_operation (p *ClaireOperation ,xt1 *ClaireType ,xt2 *ClaireType ) *ClaireType  { 
    // procedure body with s = type 
var Result *ClaireType  
    if (C_set.Id() == xt1.Isa.Id()) { 
      { var g0219 *ClaireSet   = ToSet(xt1.Id())
        if (Equal(g0219.Id(),CEMPTY.Id()) != CTRUE) { 
          Result = F_abstract_type_operation(p,F_abstract_type_set(g0219),xt2)
          } else {
          Result = ToType(g0219.Id())
          } 
        } 
      }  else if (xt1.Isa.IsIn(C_Interval) == CTRUE) { 
      { var g0220 *ClaireInterval   = To_Interval(xt1.Id())
        if (xt2.Isa.IsIn(C_Interval) == CTRUE) { 
          { var g0221 *ClaireInterval   = To_Interval(xt2.Id())
            if (p.Id() == C__plus.Id()) { 
              Result = F__dot_dot_integer((g0220.Arg1+g0221.Arg1),(g0220.Arg2+g0221.Arg2))
              }  else if (p.Id() == C__dash.Id()) { 
              Result = F__dot_dot_integer((g0220.Arg1-g0221.Arg2),(g0220.Arg2-g0221.Arg1))
              } else {
              Result = ToType(C_integer.Id())
              } 
            } 
          }  else if (C_set.Id() == xt2.Isa.Id()) { 
          { var g0222 *ClaireSet   = ToSet(xt2.Id())
            if (Equal(g0222.Id(),CEMPTY.Id()) != CTRUE) { 
              Result = F_abstract_type_operation(p,ToType(g0220.Id()),F_abstract_type_set(g0222))
              } else {
              Result = ToType(g0222.Id())
              } 
            } 
          }  else if (xt2.Isa.IsIn(C_Union) == CTRUE) { 
          { var g0223 *ClaireUnion   = To_Union(xt2.Id())
            Result = F_U_type(F_abstract_type_operation(p,ToType(g0220.Id()),g0223.T1),F_abstract_type_operation(p,ToType(g0220.Id()),g0223.T2))
            } 
          } else {
          Result = ToType(C_integer.Id())
          } 
        } 
      }  else if (xt1.Isa.IsIn(C_Union) == CTRUE) { 
      { var g0225 *ClaireUnion   = To_Union(xt1.Id())
        Result = F_U_type(F_abstract_type_operation(p,g0225.T1,xt2),F_abstract_type_operation(p,g0225.T2,xt2))
        } 
      } else {
      Result = ToType(C_integer.Id())
      } 
    return Result} 
  
// The EID go function for: abstract_type @ operation (throw: false) 
func E_abstract_type_operation (p EID,xt1 EID,xt2 EID) EID { 
    return EID{F_abstract_type_operation(ToOperation(OBJ(p)),ToType(OBJ(xt1)),ToType(OBJ(xt2)) ).Id(),0}} 
  
// we create some types that we need
// a useful second ortder type
/* {1} The go function for: first_arg_type(x:type,y:type) [status=0] */
func F_first_arg_type_type (x *ClaireType ,y *ClaireType ) *ClaireType  { 
    return  x
    } 
  
// The EID go function for: first_arg_type @ list<type_expression>(type, type) (throw: false) 
func E_first_arg_type_type (x EID,y EID) EID { 
    return EID{F_first_arg_type_type(ToType(OBJ(x)),ToType(OBJ(y)) ).Id(),0}} 
  
/* {1} The go function for: first_arg_type(x:type,y:type,z:type) [status=0] */
func F_first_arg_type_type2 (x *ClaireType ,y *ClaireType ,z *ClaireType ) *ClaireType  { 
    return  x
    } 
  
// The EID go function for: first_arg_type @ list<type_expression>(type, type, type) (throw: false) 
func E_first_arg_type_type2 (x EID,y EID,z EID) EID { 
    return EID{F_first_arg_type_type2(ToType(OBJ(x)),ToType(OBJ(y)),ToType(OBJ(z)) ).Id(),0}} 
  
/* {1} The go function for: second_arg_type(x:type,y:type) [status=0] */
func F_second_arg_type_type (x *ClaireType ,y *ClaireType ) *ClaireType  { 
    return  y
    } 
  
// The EID go function for: second_arg_type @ type (throw: false) 
func E_second_arg_type_type (x EID,y EID) EID { 
    return EID{F_second_arg_type_type(ToType(OBJ(x)),ToType(OBJ(y)) ).Id(),0}} 
  
/* {1} The go function for: meet_arg_types(x:type,y:type) [status=0] */
func F_meet_arg_types_type (x *ClaireType ,y *ClaireType ) *ClaireType  { 
    return  F_U_type(x,y)
    } 
  
// The EID go function for: meet_arg_types @ type (throw: false) 
func E_meet_arg_types_type (x EID,y EID) EID { 
    return EID{F_meet_arg_types_type(ToType(OBJ(x)),ToType(OBJ(y)) ).Id(),0}} 
  
/* {1} The go function for: first_member_type(x:type,y:type) [status=0] */
func F_first_member_type_type (x *ClaireType ,y *ClaireType ) *ClaireType  { 
    return  F_member_type(x)
    } 
  
// The EID go function for: first_member_type @ type (throw: false) 
func E_first_member_type_type (x EID,y EID) EID { 
    return EID{F_first_member_type_type(ToType(OBJ(x)),ToType(OBJ(y)) ).Id(),0}} 
  
// v3.3.10
// nth@bag (list / set) is now in Kernel (CLAIRE4)
/* {1} The go function for: nth_arg_type(x:type,y:type) [status=1] */
func F_Core_nth_arg_type_type (x *ClaireType ,y *ClaireType ) EID { 
    var Result EID 
    if ((C_tuple.Id() == x.Isa.Id()) && 
        (F_unique_ask_type(y) == CTRUE)) { 
      { var arg_1 *ClaireAny  
        _ = arg_1
        var try_2 EID 
        /*g_try(v2:"try_2",loop:false) */
        try_2 = F_the_type(y)
        /* ERROR PROTECTION INSERTED (arg_1-Result) */
        if ErrorIn(try_2) {Result = try_2
        } else {
        arg_1 = ANY(try_2)
        Result = F_CALL(C_nth,ARGS(EID{x.Id(),0},arg_1.ToEID()))
        }
        } 
      } else {
      Result = EID{F_member_type(x).Id(),0}
      } 
    return Result} 
  
// The EID go function for: nth_arg_type @ type (throw: true) 
func E_Core_nth_arg_type_type (x EID,y EID) EID { 
    return F_Core_nth_arg_type_type(ToType(OBJ(x)),ToType(OBJ(y)) )} 
  
// we place here all methods that require second order types !!!!
/* {1} The go function for: nth_get(a:array,n:integer) [status=0] */
func F_nth_get_array (a *ClaireList ,n int) *ClaireAny  { 
    return  ToList(a.Id()).At(n-1)
    } 
  
// The EID go function for: nth_get @ array (throw: false) 
func E_nth_get_array (a EID,n EID) EID { 
    return F_nth_get_array(ToArray(OBJ(a)),INT(n) ).ToEID()} 
  
/* {1} The go function for: nth_get_array_type */
func F_nth_get_array_type (a *ClaireType ,n *ClaireType ) EID { 
    var Result EID 
    Result = EID{F_member_type(a).Id(),0}
    return Result} 
  
  
// The dual EID go function for: "nth_get_array_type" 
func E_nth_get_array_type (a EID,n EID) EID { 
    return F_nth_get_array_type(ToType(OBJ(a)),ToType(OBJ(n)))} 
  
// managed by cross-compiler ?
/* {1} The go function for: nth(self:array,x:integer) [status=1] */
func F_nth_array (self *ClaireList ,x int) EID { 
    var Result EID 
    if ((x > 0) && 
        (x <= self.Length())) { 
      Result = ToList(self.Id()).At(x-1).ToEID()
      } else {
      Result = ToException(C_general_error.Make(MakeString("[180] nth[~S] out of scope for ~S").Id(),MakeConstantList(MakeInteger(x).Id(),self.Id()).Id())).Close()
      } 
    return Result} 
  
// The EID go function for: nth @ array (throw: true) 
func E_nth_array (self EID,x EID) EID { 
    return F_nth_array(ToArray(OBJ(self)),INT(x) )} 
  
/* {1} The go function for: nth_array_type */
func F_nth_array_type (self *ClaireType ,x *ClaireType ) EID { 
    var Result EID 
    Result = EID{F_member_type(self).Id(),0}
    return Result} 
  
  
// The dual EID go function for: "nth_array_type" 
func E_nth_array_type (self EID,x EID) EID { 
    return F_nth_array_type(ToType(OBJ(self)),ToType(OBJ(x)))} 
  
/* {1} The go function for: make_array_integer_type */
func F_make_array_integer_type (i *ClaireType ,t *ClaireType ,v *ClaireType ) EID { 
    var Result EID 
    if (F_unique_ask_type(t) == CTRUE) { 
      { var arg_1 *ClaireAny  
        _ = arg_1
        var try_2 EID 
        /*g_try(v2:"try_2",loop:false) */
        try_2 = F_the_type(t)
        /* ERROR PROTECTION INSERTED (arg_1-Result) */
        if ErrorIn(try_2) {Result = try_2
        } else {
        arg_1 = ANY(try_2)
        Result = EID{F_nth_type(ToType(arg_1)).Id(),0}
        }
        } 
      } else {
      Result = EID{C_array.Id(),0}
      } 
    return Result} 
  
  
// The dual EID go function for: "make_array_integer_type" 
func E_make_array_integer_type (i EID,t EID,v EID) EID { 
    return F_make_array_integer_type(ToType(OBJ(i)),ToType(OBJ(t)),ToType(OBJ(v)))} 
  
/* {1} The go function for: make_list(n:integer,t:type,x:any) [status=0] */
func F_make_list_integer2 (n int,t *ClaireType ,x *ClaireAny ) *ClaireList  { 
    return  ToList(F_make_list_integer(n,x).Cast_I(t).Id())
    } 
  
// The EID go function for: make_list @ list<type_expression>(integer, type, any) (throw: false) 
func E_make_list_integer2 (n EID,t EID,x EID) EID { 
    return EID{F_make_list_integer2(INT(n),ToType(OBJ(t)),ANY(x) ).Id(),0}} 
  
/* {1} The go function for: make_list_integer2_type */
func F_make_list_integer2_type (n *ClaireType ,t *ClaireType ,x *ClaireType ) EID { 
    var Result EID 
    if (F_unique_ask_type(t) == CTRUE) { 
      { var arg_1 *ClaireAny  
        _ = arg_1
        var try_2 EID 
        /*g_try(v2:"try_2",loop:false) */
        try_2 = F_the_type(t)
        /* ERROR PROTECTION INSERTED (arg_1-Result) */
        if ErrorIn(try_2) {Result = try_2
        } else {
        arg_1 = ANY(try_2)
        Result = EID{F_nth_class1(C_list,ToType(arg_1)).Id(),0}
        }
        } 
      } else {
      Result = EID{C_list.Id(),0}
      } 
    return Result} 
  
  
// The dual EID go function for: "make_list_integer2_type" 
func E_make_list_integer2_type (n EID,t EID,x EID) EID { 
    return F_make_list_integer2_type(ToType(OBJ(n)),ToType(OBJ(t)),ToType(OBJ(x)))} 
  
/* {1} The go function for: make_set(self:array[of:(any)]) [status=0] */
func F_make_set_array (self *ClaireList ) *ClaireSet  { 
    return  F_list_I_array(self).Set_I()
    } 
  
// The EID go function for: make_set @ array (throw: false) 
func E_make_set_array (self EID) EID { 
    return EID{F_make_set_array(ToArray(OBJ(self)) ).Id(),0}} 
  
/* {1} The go function for: make_set_array_type */
func F_make_set_array_type (self *ClaireType ) EID { 
    var Result EID 
    if (F_member_type(self.At(C_of)).Id() == C_any.Id()) { 
      Result = EID{C_set.Id(),0}
      } else {
      Result = EID{F_nth_class2(C_set,MakeList(ToType(C_any.Id()),C_of.Id()),MakeConstantList(MakeConstantSet(F_member_type(self.At(C_of)).Id()).Id())).Id(),0}
      } 
    return Result} 
  
  
// The dual EID go function for: "make_set_array_type" 
func E_make_set_array_type (self EID) EID { 
    return F_make_set_array_type(ToType(OBJ(self)))} 
  
// these four functions are defined in Core with Kernel functions because we want to
// add second order types
/* {1} The go function for: list_I_array_type */
func F_list_I_array_type (a *ClaireType ) EID { 
    var Result EID 
    if (F_member_type(a.At(C_of)).Id() == C_any.Id()) { 
      Result = EID{C_list.Id(),0}
      } else {
      Result = EID{F_nth_class2(C_list,MakeList(ToType(C_any.Id()),C_of.Id()),MakeConstantList(MakeConstantSet(F_member_type(a.At(C_of)).Id()).Id())).Id(),0}
      } 
    return Result} 
  
  
// The dual EID go function for: "list_I_array_type" 
func E_list_I_array_type (a EID) EID { 
    return F_list_I_array_type(ToType(OBJ(a)))} 
  
/* {1} The go function for: array_I_list_type */
func F_array_I_list_type (a *ClaireType ) EID { 
    var Result EID 
    if (F_member_type(a.At(C_of)).Id() == C_any.Id()) { 
      Result = EID{C_array.Id(),0}
      } else {
      Result = EID{F_nth_class2(C_array,MakeList(ToType(C_any.Id()),C_of.Id()),MakeConstantList(MakeConstantSet(F_member_type(a.At(C_of)).Id()).Id())).Id(),0}
      } 
    return Result} 
  
  
// The dual EID go function for: "array_I_list_type" 
func E_array_I_list_type (a EID) EID { 
    return F_array_I_list_type(ToType(OBJ(a)))} 
  
// v3.0.72
/* {1} The go function for: set_I_list_type */
func F_set_I_list_type (l *ClaireType ) EID { 
    var Result EID 
    if (F_member_type(l.At(C_of)).Id() == C_any.Id()) { 
      Result = EID{C_set.Id(),0}
      } else {
      Result = EID{F_nth_class2(C_set,MakeList(ToType(C_any.Id()),C_of.Id()),MakeConstantList(MakeConstantSet(F_member_type(l.At(C_of)).Id()).Id())).Id(),0}
      } 
    return Result} 
  
  
// The dual EID go function for: "set_I_list_type" 
func E_set_I_list_type (l EID) EID { 
    return F_set_I_list_type(ToType(OBJ(l)))} 
  
/* {1} The go function for: list_I_set_type */
func F_list_I_set_type (l *ClaireType ) EID { 
    var Result EID 
    if (F_member_type(l.At(C_of)).Id() == C_any.Id()) { 
      Result = EID{C_list.Id(),0}
      } else {
      Result = EID{F_nth_class2(C_list,MakeList(ToType(C_any.Id()),C_of.Id()),MakeConstantList(MakeConstantSet(F_member_type(l.At(C_of)).Id()).Id())).Id(),0}
      } 
    return Result} 
  
  
// The dual EID go function for: "list_I_set_type" 
func E_list_I_set_type (l EID) EID { 
    return F_list_I_set_type(ToType(OBJ(l)))} 
  
// get the type from class if a constant
/* {1} The go function for: thing_type_class(x:type) [status=0] */
func F_Core_thing_type_class_type (x *ClaireType ) *ClaireType  { 
    return  F_glb_class(C_thing,F_member_type(x))
    } 
  
// The EID go function for: thing_type_class @ type (throw: false) 
func E_Core_thing_type_class_type (x EID) EID { 
    return EID{F_Core_thing_type_class_type(ToType(OBJ(x)) ).Id(),0}} 
  
/* {1} The go function for: object_type_class(x:type) [status=0] */
func F_Core_object_type_class_type (x *ClaireType ) *ClaireType  { 
    return  F_glb_class(C_object,F_member_type(x))
    } 
  
// The EID go function for: object_type_class @ type (throw: false) 
func E_Core_object_type_class_type (x EID) EID { 
    return EID{F_Core_object_type_class_type(ToType(OBJ(x)) ).Id(),0}} 
  
// new in v3.0.60 : second-order type for copy